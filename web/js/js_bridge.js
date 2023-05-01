async function pauseExecResume(messageToEcho, execMe) {
    this.echo(messageToEcho)
    this.pause()
    let result = await execMe()
    this.resume()
    return result
}

async function listOrganisations(token) {
    let orgs = await pauseExecResume.call(this,
        "Listing orgs...",
        async () => JSON.parse(await listOrgs(token))
    )
    for (let i = 0; i < orgs.length; i++) {
        this.echo(`OrganisationId: ${orgs[i].OrganisationId}`)
    }
}

async function createOrganisation(token, organisationId, organisationName) {
    let result = await pauseExecResume.call(this,
        "Creating organisation. Please wait, this can take up to a minute",
        async () => await createOrg(token, organisationId, organisationName)
    )

    this.echo("Done! Organisation created with Id: " + organisationId)
}

async function deleteApplication(token, organisationId, applicationId) {
    let result = await pauseExecResume.call(this,
        "Removing application. Please wait, this can take up to a minute",
        async () => await deleteApp(token, organisationId, applicationId)
    )
    this.echo(result)
}

async function createApplication(token, organisationId, applicationName) {
    let result = await pauseExecResume.call(this,
        "Creating application. Please wait, this can take up to a minute",
        async () => JSON.parse(await createApp(token, organisationId, applicationName))
    )

    this.echo("Done! Application created with Id: " + result.ApplicationId)
}

async function listApplications(token, organisationId, showApiKey) {
    let apps = await pauseExecResume.call(this,
        "Listing apps...",
        async () => JSON.parse(await listApps(token, organisationId))
    )
    for (let i = 0; i < apps.length; i++) {
        this.echo("*****")
        this.echo(`ApplicationId: ${apps[i].ApplicationId}`)
        this.echo(`ApplicationName: ${apps[i].ApplicationName}`)

        if (showApiKey) {
            this.echo(`ApiKey: ${apps[i].ApiKey}`)
        }
        this.echo("-----")
    }
}

async function deleteOrganisation(token, organisationId) {
    let result = await pauseExecResume.call(this,
        "Removing organisation. Please wait, this can take up to a minute",
        async () => await deleteOrg(token, organisationId)
    )
    this.echo(result)
}