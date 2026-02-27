# Branching Strategy

These are the branches that will exist in this repo and their purposes:

## `main` Branch

* Used for production deployments.
* Builds must always be in a stable, working state and able to deploy (i.e. `main` should never contain broken code).
* All production deployments should run through a CI/CD pipeline that runs automated tests before deployment. If all tests pass, then deployment continues.
* After a production deployment is live, all manual and automated tests for the live production app should occur.

<br>

## `staging` Branch

* The `staging` branch and environment should be an exact clone of the `main` branch and environment.
* Used as the integration branch for all feature/development branches (i.e. all feature/development branches will branch off of `staging`).
* Staging deployments occur from `staging`.
* Only after all the manual/automated tests pass should you merge the `staging` branch into `main` and test a production deployment.
* All `staging` deployments should run through a CI/CD pipeline that runs automated tests before deployment. If all tests pass, then deployment continues.
* After a `staging` deployment is live, all manual and automated tests for the live `staging` app should occur.

<br>

## Feature/Development Branches

* All feature/development branches will branch off of `staging`.
* When a feature branch is complete you should test a preview deployment and then a `staging` deployment.
    * The `staging` branch should be merged into the feature branch and the feature branch should be tested within it's own branch/environment.
    * Only after all the manual/automated tests pass should you merge the feature branch into `staging` and test a `staging` deployment.

<br>

## Branch Prefixes

* Prefixes like `feature/`, `bugfix/`, `hotfix/`, or `chore/` immediately convey the purpose of the branch without needing to examine its contents or commit history.
* Many CI/CD tools and other automation services can be configured to recognize specific branch prefixes and trigger corresponding workflows, such as running specific tests for `bugfix/` branches or deploying `release/` branches.

### Ideas for Prefixes & Structure

* `feature/`: For developing new features. Example: `feature/user-authentication`
* `bugfix/`: For fixing bugs. Example: `bugfix/login-timeout`
* `hotfix/`: For urgent fixes on production. Example: `hotfix/security-vulnerability`
* `chore/`: For maintenance tasks or minor changes that are not features or bug fixes. Example: `chore/update-dependencies`
* `release/`: For preparing a new release. Example: `release/v1.0.0`

### Additional Considerations

* Issue/Ticket Numbers: Including the relevant issue or ticket number from a project management tool (e.g., Jira, Trello) can further enhance traceability. Example: `feature/JIRA-123-new-login-system`
* Kebab Case: Use hyphens to separate words in the descriptive part of the branch name for improved readability (e.g., `user-authentication` instead of `userAuthentication`).
* Consistency: The most crucial aspect is to establish and consistently follow a naming convention within your team.
