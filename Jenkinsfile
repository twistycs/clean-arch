pipeline {
    agent any
    stages {
        stage('Cloning Git') {
        steps {
            git([url: 'https://github.com/twistycs/clean-arch.git', branch: 'dev', credentialsId: 'wittharit-github-user-token'])
        }

        stage('Test Echo') {
        steps {
            Echo 'Test Echo12334455556664444'
        }
    }
}