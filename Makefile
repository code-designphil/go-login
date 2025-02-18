deploy:
	cd lambda && make build
	cd ..
	cdk deploy
