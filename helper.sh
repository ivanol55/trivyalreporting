#!/bin/bash

# Runs when no option is provided or help is requested
if [ -z "$1" ] || [ "$1" == "-h" ] || [ "$1" == "--help" ] || [ "$1" == "help" ]; then
    echo "# Please provide one of the following to start: "
    echo "- 'start' to run the reports web server"
    echo "- 'report' to run a report and put it on /latest/, includes ec2 and s3 services"
    echo "- 'multireport' to run an on-demand report on multiple services, with structure /YYYY-MM-DD-HH-MM/[service] (this demo creates /[datetime]/s3 and /[datetime]/ec2)"
    echo "- 'stop' to stop the web server environment"
    exit 1
fi

case $1 in
    # Start the reports webserver
    start)
        echo "Starting reports webserver..."
        docker compose up -d
    ;;

    # Run a report and put it into /latest/
    report)
        echo "Running a general report for /latest/..."
        docker run \
            -v /home/$USER/.aws/:/root/.aws/ \
            -v $(pwd)/src/:/src/ \
            -v $(pwd)/webfiles/:/webfiles/ \
            aquasec/trivy aws \
            --region eu-central-1 \
            --service ec2 \
            --service s3 \
            --format template \
            --template @/src/template.tpl \
            -o /webfiles/reports/latest/index.html
        echo ""
        echo "If the reports webserver is running, you can access this report at http://localhost:8080/reports/latest/"
        echo "You can start the reports webserver by running this same script with the 'start' flag: \`./helper.sh start\`"
    ;;

    # Run an on-demand report for each listed service
    multireport)
        echo "Running an on-demand multireport..."
        foldername=$(date +"%Y-%m-%d-%H-%M")
        mkdir ./webfiles/reports/$foldername
        services=("ec2" "s3")
        for service in ${services[@]}; do
            mkdir ./webfiles/reports/$foldername/$service
            docker run \
                -v /home/$USER/.aws/:/root/.aws/ \
                -v $(pwd)/src/:/src/ \
                -v $(pwd)/webfiles/:/webfiles/ \
                aquasec/trivy aws \
                --region eu-central-1 \
                --service $service \
                --format template \
                --template @/src/template.tpl \
                -o /webfiles/reports/$foldername/$service/index.html            
        done
        echo "Your on-demand report for each service is available at http://localhost:8080/reports/$foldername/ with each service under it"
        echo "You can start the reports webserver by running this same script with the 'start' flag: \`./helper.sh start\`"
    ;;

    # Start the reports webserver
    start)
        echo "Stopping reports webserver..."
        docker compose down
    ;;

    # catch-all case
    *)
        echo "Unknown option. Please provide one of the supported options (see by running the script with no provided inputs or with the 'help' flag):"
        echo ""
        echo -e "\t./reporting.sh help"
        exit 1
    ;;

esac
