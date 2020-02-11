
# GoLang Currency API

GoLang Currency API is a microservice for getting currency exchange rates. The API returns the value of GBP and USD currency against EUR.

## Installation

Use the makefile instructions to build a docker image and create a running container ready to serve requests. From the project's root path, run the following commands:

```bash
make docker_build_image
make docker_run_container
```
Or simply run 

```bash
make all
```

## Usage
The API expose one endpoint that accepts a currency parameter
- currency: the currency code to be requested. Only USD and GBP are supported. If no currency is provided, the system uses GBP as default

## ToDo
The API is far from being completed, several items needs to be added:
1. Unit Tests, Acceptance Tests
2. Support for multiple currencies
3. Naive recommendation to exchange money based on the previous week historical data
4. CI/CD pipeline
	4.1 Download dependencies
	4.2 Build code
	4.3 Run Tests
	4.4 Deploy to environments

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
