[CmdletBinding()]
param (
    [Parameter(Mandatory=$false)]
    [string]$version="v0.0.3"
)
go env -w GOPRIVATE=github.com/Rukas/*
go get -u github.com/Rukas/librenms-go-client@$version