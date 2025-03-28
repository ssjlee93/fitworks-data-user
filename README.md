# fitworks-data-user
data service for FitWorks dedicated to user service

## Notes

### Repository pattern
a data retrieval pattern based on business logic.  
DAO is meant to access DB.  
Repository pattern retrieves curtailed to the business logic.  

This data service only retrieves data.  
logic should be handled on a different microservice.  
hence, I did nothing but toss the results of the DAOs.

## dtos vs entities
entities makes more sense.  
but for now, keeping it DTOs since they are the same right now.  
