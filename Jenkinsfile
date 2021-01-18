pipeline {
    agent any

    stages {
        stage('Test') {
            steps {
                echo 'build/build-bin.sh'
            }
        }
        stage('Test1') {
            steps {
                echo 'build/run-tests.sh'
            }
        }
    }
}