pipeline {
    stages {
        stage('Test') {
            
            steps {
                sh 'go test ./... '                
            }
        }
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        
        stage('Code Analysis') {
            steps {
                sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
                sh 'golangci-lint run'
            }
        }        
    }
}