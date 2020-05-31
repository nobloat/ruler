**CONSTRUCTION SITE - WORK IN PROGRESS - MIGHT END UP LIKE BER**

# ruler
Tool to track dependncies of software-projects

## Metrics 
- **Lines of source** and **Number of files** per category
  - Test
  - Examples
  - Functionality
- Resulting optimized binary size of the functionality including runtime-dependencies
- Total install-size of dev-dependencies in Bytes (on top of alpine-linux docker image) -> Pushed to docker registry
- Total install-size of external-dependencies -> Use docker image size if available
- Number of service dependencies (size cannot be calculated)

- Add file to repo -> Dockerfile.ruler
- Use resulting docker image sizes for dev-dependencies and total optimized binary size with runtime
- Cound lines of code / source lines based on three categories via simple tool

## Dependency types
### Dev-dependencies
- Everything a contributor needs to install in order to contribute to the project
- Typical examples: Testing frameworks, `"devDependencies"` section in `package.json`, build systems, compilers
### Runtime-dependencies
- Everything the user of a software component needs in order to use/run the component/application
### External-dependencies
- Software artifacts that do not reside in the same process (memory-address-space) of the component (e.g. databases, web-services, CDNs)
### Service-dependencies
- Software that is not available as source-code nor binary and can therefore not be measured except for its existence
### Implicit-dependencies
- e.g. operating system, are currently not measured

## Goals
- Manuals / best-practices to cont. measure and track the above mentioned metrics
- Automated tooling to retrieve these metrics
- Visualize the results for some open-source projects
- Easy to use and integrate into existing projects -> Badges based on github repo urls
- simple-go binary instead of bash script that calls docker history to export layer sizes + does line counting for the actual code.

## Non-goals
- https://dependencytrack.org/
- Building the docker images on separte infrastructure -> use what is already there.

## See also
- https://github.com/wagoodman/dive
- https://github.com/XAMPPRocky/tokei_rs Lines of Code badge
