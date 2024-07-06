package com.ppuu444.h2memdb;


import com.ppuu444.h2memdb.model.Executions;
import com.ppuu444.h2memdb.repository.ExecutionsRepository;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class H2memdbApplication {

	public static void main(String[] args) {
		System.out.println("starting cache on port: 8083");
		SpringApplication.run(H2memdbApplication.class, args);
	}

}
