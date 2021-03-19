#Log file aggregator

- goLang api for product decisioning based off testing logs
- dockerized
- Helm Chart ready to be deployed to kubernetes

###Assumptions

I made some assumptions around this to make my life easier and also provide a platform for expansion

- The logs will always be in this format
   * reference
   * product
   * datapoints about product

###Where is the duct tape

- the logic here is not great (It works it just could be a lot better)
- timestamp parsing is wack as i concat'd timezone info to get it into RFC3339 time formatting
- I would say this is not expandable right now due to logic and source logs I think to make this more expandable the input to this application would need to reworking

###How would I improve upon this application

- source logs should be updated into json formatting
- source logs shipped to s3
- this application could then pull from s3 (instead of a local file)
- logs can then be mapped into structs and the logic can be cleaned up to support additional product types and datapoints

###Improvements I have made

- output format includes time window of testing
- seperation of prodcuts in output formatting
- expandable output object



