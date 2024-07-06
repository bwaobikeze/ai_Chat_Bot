package com.ppuu444.h2memdb.controller;

import com.ppuu444.h2memdb.model.Executions;
import com.ppuu444.h2memdb.repository.ExecutionsRepository;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Iterator;
import java.util.Objects;

@RestController
@RequestMapping("/")
public class ExecutionsController {
    private final ExecutionsRepository repository;

    public ExecutionsController(ExecutionsRepository repository) {
        this.repository = repository;
    }

    @PostMapping
    public ResponseEntity<String> postController(@RequestBody ExecutionObject body) {
        System.out.println("==updating execution record==");
        Integer executionId = null;
        try{
            Iterator<Executions> results = repository.findAll().iterator();
            Integer hashedQuery = body.query.hashCode();

            while(results.hasNext()){
                Executions currentResult = results.next();
                if(Objects.equals(currentResult.hash(), hashedQuery)) {
                    executionId = currentResult.id();
                    break;
                }
            }

            repository.save( new Executions(executionId, body.query.hashCode(), body.result));
        } catch (Exception e) {
            System.out.println("error updating execution record");
            e.printStackTrace();
            executionId = -1;
        }
        if (executionId != null && executionId < 0){
            return ResponseEntity.ok("unable to update query");
        }
        return ResponseEntity.ok("updated query successfully");
    }

    @GetMapping
    public ResponseEntity<String> getController(@RequestBody RequestObject body){
        System.out.println("==getting query result==");
        String response = "";
        try {
            Iterator<Executions> results = repository.findAll().iterator();
            Integer hashedQuery = body.query.hashCode();

            while(results.hasNext()){
                Executions currentResult = results.next();
                if(Objects.equals(currentResult.hash(), hashedQuery)) {
                    response = currentResult.result();
                    break;
                }
            }
        } catch (Exception e) {
            System.out.println("error getting query result");
            e.printStackTrace();
        }

        if(response.length() > 0) {
            System.out.println("found query");
        } else {
            System.out.println("no query found");
        }

        return ResponseEntity.ok(response);
    }

    @DeleteMapping
    public ResponseEntity<String> deleteAll() {
        System.out.println("==clearing db==");
        try{
            repository.deleteAll();
            return ResponseEntity.ok("db cleared successfully");
        } catch (Exception e) {
            System.out.println("error clearing db");
            e.printStackTrace();
        }
        return ResponseEntity.ok("unable to clear db");
    }
}
