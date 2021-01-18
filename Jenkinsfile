pipeline {
    agent any

    stages {
        stage('Test1') {
            steps {
                echo 'build/build-bin.sh'
            }
        }
        stage('Test2') {
            steps {
                echo 'build/run-tests.sh'
            }
        }
    }
}