#!groovy
pipeline {
        agent any
        parameters {
		
		    //1
		   	choice(name: 'black_duck_server', choices: ['https://blackduck.com/api/','https://blackduck-test.com/api/'], description: 'Blackduck server')
			//2
			string(name: 'black_duck_api_key',defaultValue: '', description: 'API Key is of Dinesh Pls provide your own API key from blackduck account while running the job. \nUse test server api key then use this "gxYy1iNjk0LTJjNzRmZWVkZTYxNQ==\n prod server is YjYzNTlmZjIxLTMwZjMt4MmUyNQ"', trim: true)
           
		    //3
		    string(name: 'projectversionurl',defaultValue: '', description: 'full projectversionurl', trim: true)
			//4
		    string(name: 'codelocationsurl',defaultValue: '', description: 'full codelocationsurl', trim: true)
			//5
		    string(name: 'matchtocomponentversionurl',defaultValue: '', description: 'matchtocomponentversionurl', trim: true)
			//6
		    string(name: 'matchString',defaultValue: '', description: 'matchString', trim: true)
			//7
		    string(name: 'entriestoresolve',defaultValue: '10', description: 'entriestoresolve', trim: true)
	
        }
    
       stages {
            stage('Trigger bulk match') {
                steps {
                    script {
					
							sh """
							   #!/bin/bash 
							    dpkg --print-architecture
							    uname -a
							    rm -rv hubsm*
								wget https://github.com/dineshr93/blackduckAPI/releases/download/0.0.2/hubsm-amd64-linuxbuilt -O hubsm
								chmod +x hubsm
								ls -l
								./hubsm ${params.black_duck_server} ${params.black_duck_api_key} ${params.projectversionurl} ${params.codelocationsurl} ${params.matchtocomponentversionurl} ${params.matchString} ${params.entriestoresolve}

							"""
							echo "======================Completed=========================="
                        } //script
                } //step
            } //stage
       } //stages
    }//pipeline