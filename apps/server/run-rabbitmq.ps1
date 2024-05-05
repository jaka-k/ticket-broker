$containerName = "rabbitmq-dev"

# Check if Docker is installed
if (!(Get-Command docker -ErrorAction SilentlyContinue)) {
    Write-Host "Error: Docker is not installed. Please install Docker to run RabbitMQ." -ForegroundColor Red
    exit 1
}

# Check if Docker is running
if (-not (docker info > $null 2>&1)) {
    Write-Host "Error: Docker is not running. Please start Docker to run RabbitMQ." -ForegroundColor Red
    exit 1
}

# Check if RabbitMQ is already running
if (docker ps -q -f "name=$containerName") {
    Write-Host "RabbitMQ is already running." -ForegroundColor Green
}
else {
    # Check if the container exists but is stopped
    if (docker ps -aq -f "status=exited" -f "name=$containerName") {
        Write-Host "Starting existing RabbitMQ container..." -ForegroundColor Yellow
        docker start $containerName
    }
    else {
        Write-Host "Creating and starting new RabbitMQ container..." -ForegroundColor Yellow
        docker run -d --name $containerName -p 5672:5672 -p 15672:15672 rabbitmq:management
    }
}
