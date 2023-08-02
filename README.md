gate-zkmerkle-proof repo

## Prerequisit
You need to install  mysql, redis , kvrocks

##command
```
    ./main.go keygen     // zk key generate
    ./mian.go witness   // generate witness data
    ./mian.go prover   // generate zk proof 
    ./mian.go userproof   // generate zk proof 
    ./main.go verify cex  
    ./main.go verify user
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

