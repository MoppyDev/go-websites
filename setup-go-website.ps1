<#
.Synopsis
setup-go-website.ps1 is used to create the base folders and files used for a basic Go website

.Description
setup-go-website.ps1 created to assist in making a uniformed folder structure for a new Go based website in the current directory.

.Parameter websiteName
Sets the name for the website you want to create. This will create a base folder for the child folders and files to live in. 

.Example
setup-go-website.ps1 -websiteName "Porfolio"
#>

Param(
    [Parameter(mandatory=$true)][String]$websiteName
)

if (!(Test-Path $websiteName))
{
    try 
    {
        write-host "Creating recommended folders for your new Go website [$websiteName]"
        New-Item -ItemType Directory -Name $websiteName
        Set-Location $websiteName
        New-item -ItemType Directory -Name "cmd/web"
        New-item -ItemType Directory -Name "internal"
        New-item -ItemType Directory -Name "ui/html/pages"
        New-item -ItemType Directory -Name "ui/html/partials"
        New-item -ItemType Directory -Name "ui/static/css"
        New-item -ItemType Directory -Name "ui/static/img"
        New-item -ItemType Directory -Name "ui/static/js"

        write-host "Creating recommended starter files in cmd/web: main.go and handlers.go"
        New-Item -ItemType File -Name "cmd/web/main.go"
        New-Item -ItemType File -Name "cmd/web/handlers.go"
    }
    catch
    {
        $ErrorMessage = $_.Exception.Message
        write-error "Failed to create Go website folder structure: $ErrorMessage"
    }
}
else {
    Write-Warning "Folder $websiteName already exists.."
}