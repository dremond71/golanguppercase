pipeline {
    agent any
    tools {
        go 'go1.17.3'
    }

    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {        
        stage('Install Dependencies') {
            steps {
                echo 'Installing dependencies...'
                sh 'go version'
                sh 'go install golang.org/x/lint/golint@latest'
                sh 'go install github.com/onsi/ginkgo/ginkgo@latest' 
                sh 'go get github.com/onsi/gomega/...'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building...'
                sh 'go build'
            }
        }


        stage('Test') {
            steps {
                   withEnv(["PATH+GO=${GOPATH}/bin"]){
                        echo 'Linting...'
                        sh 'golint .'
                        echo ''
                        echo 'Running tests...'
                        sh 'ginkgo -r -v'
                   }
            }
        }
        
    }
  
}