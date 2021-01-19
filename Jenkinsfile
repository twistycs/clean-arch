pipeline {
    agent any
    stages {
        stage('Cloning Git') {
            steps {
                git([url: 'https://github.com/twistycs/clean-arch.git', branch: 'dev', credentialsId: 'wittharit-github-user-token'])
            }
        }

        stage('Test Echo') {
            steps {
                echo "Test Echo1233445555666444455555555"
            }
        }
   
    }
}