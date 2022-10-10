# What is TrivyalReporting?
`trivyalreporting` (Trivial + Trivy + Reporting) is a **set of tools and scripts** that aims to provide you with a continuous infrastructure security posture management frontend based on a **constantly generated**, **lightweight**, **quickly actionable** set of flat HTML and CSS files.

# Why do we need this?
Infrastructure security posture management is a **considerable undertaking**. Data needs to be gathered from several sources, and all this data needs to be **processed**, **aggregated** and **transformed** into easy to recognize, actionable alerts and reports. In our case, to avoid further service costs and vendor lock-in, we run these checks using [Trivy](https://github.com/aquasecurity/trivy) by AquaSecurity.

# Why create this tool if Trivy exists?
Trivy is very good at its job of investigating security violations against a set of standards, be it aws `infrastructure` (our case), `docker images`, or a github repository with mistakenly embedded `credentials`. That said, Trivy is designed to integrate with **CI/CD workflows** and terminal interfaces. It's really nice to have a way of automagically stopping insecure releases before merging a PR, but a 500 line `json` pasted in a comment is not the **human-centric** way of working. That's where `trivyalreporting` comes in.

# What does TrivyalReporting bring to the table?
The aim of this tool is to bring in the human-centric part we miss from AquaSecurity's already awesome [Trivy](https://github.com/aquasecurity/trivy):
- **Ease of use**: people outside infrastructure engineering may want to report their own fix requests, this empowers them to know how and what can be improved. Security culture across the board!
- **A friendly interface**: read up on what needs to be fixed without the need of a terminal
- **A method to the madness**: keep an order to the reports we generate
- **Information updates**: A point-in-time scan is useless in a constantly evolving environment. So we make it easy to keep the report up to date.
- **Narrow down on your target**: If you're wary of an s3 bucket you're creating, why do you need a report with instance findings? answer is you don't
- **Make it fancy**: what's an infrastructure report without some fancy colors?

# Is this a complete solution?
THis repository **doesn't try to be** a complete implementation of this solution. Instead, it's a *guidance* of what needs to be implemented. Need it on Kubernetes? make a manifest with the commands. Want to run it on lambdas? build a layer with the tooling! the base implementation is the **same idea** however it's implmented.

# Demo environment requirements
The entire development environment is dockerized and configured as code in `docker-compose.yaml` and the repository's file structure. You just need
- the `git` command-line interface
- the `aws` command-line interface
- AWS credentials for your accounts so they can be scanned (if you can `aws s3 ls` chances are it'll work)
- `docker` to run the webserver container for development
- the `docker` command-line interface plugin for `compose`
- `trivy` to run the scans
- `golang` 1.18 to run the helper script
- a `shell`, like `bash`, `sh` or `zsh`

# Instructions
1. Make sure all of the demo environment requirements are installed and available on your system
2. Clone this repository with `git clone git@github.com:ivanol55/trivyalreporting.git`
3. enter the cloned repository folder with `cd trivyalreporting`
4. build the helper script: `go build helper.go`
5. Run the testing environment! See some help with `./helper help`

# I'd like to contribute! What's left to do?
Left to do? Most of the things! Here's a list of stuff we should look into as this tool evolves, in no order of priority:
- ~~[ ] Implement the [client-server trivy architecture](https://aquasecurity.github.io/trivy/v0.17.0/modes/client-server/) so it doesn't need to re-download the database every time it runs on a serverless environment~~ discarded due to caching issues
- [ ] Index generator and updater: every time a report runs, update a "main page" index on `/`, since the s3 backend intented by default cannot list directories
- [x] CSS styling: let's make **TrivyalReporting** look a bit more **fancy** across the board
- [x] Establish a report format to make files smaller and easier to explore? Or keep them in a single file for simplicity? Open for discussion, as a large environment can output quite a report.
- ~~[ ] An option to output and pipe `events` into the desired event management platform (Read: How do we get this into Datadog)~~ discarded, seems out of scope for this tool
- [x] Priority filtering: on-demand reports may be overwhelming as a first approach, so let's get a way to output only `CRITICAL` or `HIGH` events
- [ ] Generic click-and-done github actions pipeline for custom reports
- [x] ~~Is date and time a good format for on-demand reports? Should we let users name the reports in any way?~~ the less user input to control, the better.
