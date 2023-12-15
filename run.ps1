param (
    [string]$ItemName,
    [string]$ApiKey
)

$syntaxError = "Missing arguments. Syntax is: './run -ItemName <item> -ApiKey <ApiKey>'"

if (-not $ItemName -or -not $ApiKey) {
    Write-Host $syntaxError
    exit
}

go build
Start-Process -WindowStyle hidden -FilePath The-Lebron-James-Notifier.exe -ArgumentList $ItemName, $ApiKey