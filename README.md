# dep-tracker
Tool to track dependncies of software-projects

## Metrics 
- Lines of code
- Number of files
- Resulting binary size

## Dependency types

### Dev-dependencies
- Everything a contributor needs to install in order to contribute to the project
- Typical examples: Testing frameworks, `"devDependencies"` section in `package.json`, build systems, compilers
### Runtime-dependencies
- Everything the user of a software component needs in order to use/run the component/application
### External-dependencies
- Software artifacts that do not reside in the same process (memory-address-space) of the component (e.g. databases, web-services, CDNs)

### Implicit-dependencies
- e.g. operating system, are currently out of the scope of this project.
