pipeline {
    agent any
    stages {
        stage('Cloning Git') {
        steps {
            git([url: 'https://github.com/twistycs/clean-arch.git', branch: 'dev', credentialsId: 'wittharit-github-user-token'])
        }
    }
}