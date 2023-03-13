async function pauseExecResume(messageToEcho, execMe) {
    this.echo(messageToEcho)
    this.pause()
    let result = await execMe()
    this.resume()
    return result
}

async function listOrganisations(token) {
    let orgs = await pauseExecResume.call(this,
        "Listing orgs.. One second..",
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

async function deleteOrganisation(token, organisationId) {
    let result = await pauseExecResume.call(this,
        "Removing organisation. Please wait, this can take up to a minute",
        async () => await deleteOrg(token, organisationId)
    )
    this.echo(result)
}