pipeline {
    agent any
    tools {
        go 'go1.17.3'
    }

    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go install golang.org/x/lint/golint@latest'
                sh 'go install github.com/onsi/ginkgo/ginkgo@latest' 
                sh 'go get github.com/onsi/gomega/...'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }


        stage('Test') {
            steps {
                    echo 'Running linting'
                    sh 'golint .'
                    echo 'Running test'
                    sh 'ginkgo -r -v --randomizeAllSpecs --randomizeSuites --race --trace'
            }
        }
        
    }
  
}