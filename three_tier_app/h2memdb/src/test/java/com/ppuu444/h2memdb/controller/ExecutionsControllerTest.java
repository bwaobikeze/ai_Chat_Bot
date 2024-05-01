package com.ppuu444.h2memdb.controller;

import com.ppuu444.h2memdb.model.Executions;
import com.ppuu444.h2memdb.repository.ExecutionsRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;

import java.util.Arrays;
import java.util.List;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

class ExecutionsControllerTest {

    private final ExecutionsRepository repository = Mockito.mock(ExecutionsRepository.class);
    private final ExecutionsController controller = new ExecutionsController(repository);

    private final MockMvc mockMvc = MockMvcBuilders.standaloneSetup(controller).build();

    @BeforeEach
    void setUp() {
        // Initialize the repository with test data if needed
        Executions testExecution = new Executions(1, "SELECT * FROM table".hashCode(), "Result");
        List<Executions> execList = Arrays.asList(new Executions[] {testExecution});
        when(repository.findAll()).thenReturn(execList);
    }

    @Test
    void postController_Success() throws Exception {
        // Arrange
        String requestBody = "{\"query\":\"SELECT * FROM table\",\"result\":\"Result\"}";

        // Mock repository behavior
        when(repository.save(any(Executions.class))).thenReturn(new Executions(2, "SELECT * FROM table".hashCode(), "Result"));

        // Act
        mockMvc.perform(post("/")
                        .content(requestBody)
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("updated query successfully"));
    }

    @Test
    void postController_Failure() throws Exception {
        // Arrange
        String requestBody = "{\"query\":\"SELECT * FROM table\",\"result\":\"Result\"}";

        // Mock repository behavior to simulate an exception
        when(repository.save(any(Executions.class))).thenThrow(new RuntimeException("Simulated exception"));

        // Act
        mockMvc.perform(post("/")
                        .content(requestBody)
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("unable to update query"));
    }

    @Test
    void getController_Success() throws Exception {
        // Arrange
        String requestBody = "{\"query\":\"SELECT * FROM table\"}";

        // Act
        mockMvc.perform(get("/")
                        .content(requestBody)
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("Result"));
    }

    @Test
    void getController_NoResultFound() throws Exception {
        // Arrange
        String requestBody = "{\"query\":\"INVALID_QUERY\"}";

        // Act
        mockMvc.perform(get("/")
                        .content(requestBody)
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string(""));
    }

    @Test
    void deleteAll_Success() throws Exception {
        // Act
        mockMvc.perform(delete("/")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("db cleared successfully"));
    }
}
