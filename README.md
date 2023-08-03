gate-zkmerkle-proof repo

## Prerequisit
You need to install  mysql, redis , kvrocks

## install
```
    make build-darwin   // compile on mac
```

## command
```
    ./main keygen     // zk key generate
    ./mian witness   // generate witness data
    ./mian prover   // generate zk proof 
    ./mian userproof   // generate zk proof 
    ./main verify cex  
    ./main verify user
```

## directory structure
```
-circuit   
-client    
-config    
-global    
-service
    --keygen_service  
    --prover_service  
    --tool_service    
    --userproff_service  
    --verify_service  
    --witness_service 
-utils  
```

