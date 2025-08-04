pipeline {
	agent {
		docker {
			image 'golang:1.23'
        }
	}

	stages {
		stage('Checkout') {
			steps {
				git branch: 'main', url: 'https://github.com/IGA/GoApiTest.git'
			}
		}
		stage('Go Mod') {
			steps {
				sh 'go mod tidy'
			}
		}
		stage('Build') {
			steps {
				sh 'go build -v ./...'
			}
		}
		stage('Test') {
			steps {
				sh 'go test -v ./...'
			}
		}
	}

	post {
		success {
			echo 'Build ve test başarılı!'
		}
		failure {
			echo 'Pipeline başarısız oldu.'
		}
	}
}