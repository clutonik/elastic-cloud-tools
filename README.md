# elastic-cloud-tools
Repository with custom tool to work with Elastic Cloud Enterprise

### Description
This repository contains code to work with Elastic Cloud Enterprise. The code
in this repository is written in golang and we have tried to cover the code
with as much test cases as we could. Feel free to open a pull request to
add new functionality.

### Tools
1. `eceRolesToken.go` : This file is used to interact with ECE during the
installation process to get Enrollment Tokens while adding secondary nodes
to an existing ECE Platform.

    `Environment Variables:`
    - Test code for this tool needs below environment variables to work:
      - `ECE_COORDINATOR_HOST`
      - `ECE_USER`
      - `ECE_PASSWORD` 
    - This tool also accepts command line flags and can be used as:
      - `eceRolesToken -coordinator-host (coordinator-hostname) 
      -ece-password (ece-root-password) 
      -ece-user (ece-root-username) -runner-role (allocator/proxy/director)`