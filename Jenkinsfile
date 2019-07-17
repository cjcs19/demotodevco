pipeline {
    agent any
    tools {
        go 'go-1.12'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Test') {
            
            steps {
                sh 'go test ./... '                
            }
        }    
        
        stage('Code Analysis') {
            steps {
                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
                sh 'golangci-lint run'
            }
        }

        stage('Compile') {
            steps {
                sh 'go build'
            }
        }        
    }
}