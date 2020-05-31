**CONSTRUCTION SITE - WORK IN PROGRESS - MIGHT END UP LIKE BER**

# ruler
Tool to track dependncies of software-projects

## Metrics 
- Lines of code
  - Test
  - Examples
  - Functionality
  - Runtime dependendencies
- Number of files
  - Test
  - Examples
  - Functionality
  - Runtime dependencies
- Resulting optimized binary size of the functionality including runtime-dependencies
- Total install-size of dev-dependencies in Bytes (may vary on each OS, suggestion to use alpine-linux as base)
- Total install-size of external-dependencies
- Number of service dependencies (size cannot be calculated)

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
- Easy to use and integrate into existing projects

## Non-goals
- https://dependencytrack.org/

## See also
- https://github.com/wagoodman/dive
