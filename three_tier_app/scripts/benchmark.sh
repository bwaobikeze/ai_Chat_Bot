#!/bin/bash

# Set the Cypher queries
cypher_queries=(
    "MATCH (n) RETURN n"
    "MATCH (a)-[:HAS]->(b:Skill) WHERE b.type = 'Database Management' RETURN a, b"
    "CREATE (n:Person {name: 'Daniel', age: 35, hobby: 'Chess'})"
    "MATCH (a)-[:VISITED]->(b:City) WHERE b.name = 'Tokyo' RETURN a, b"
    "MATCH (n)-[:FOLLOWS]->(m:Person) WHERE m.name = 'Sarah' RETURN n, m"
    "CREATE (n:Movie {title: 'The Matrix Reloaded', year: 2003, genre: 'Sci-Fi'})"
    "MATCH (a)-[:FRIEND]->(b) WHERE a.name = 'Alice' RETURN b"
    "CREATE (n:Person {name: 'John', age: 30})"
    "MATCH (n)-[:LIKES]->(m) RETURN n, m"
    "MATCH (a)-[:FOLLOWS]->(b) RETURN a, b"
    "CREATE (n:Movie {title: 'Inception', year: 2010})"
    "MATCH (n:Person)-[:WORKS_AT]->(c:Company) WHERE n.age > 25 RETURN n, c"
    "MATCH (a)-[:OWNS]->(b:Car) WHERE b.year > 2015 RETURN a, b"
    "MATCH (n)-[:FRIEND]->(m) WHERE n.name = 'Bob' RETURN m"
    "CREATE (n:Person {name: 'Emma', age: 28})"
    "MATCH (a)-[:VISITED]->(b:City) WHERE b.name = 'Paris' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Book) WHERE m.genre = 'Science Fiction' RETURN n, m"
    "MATCH (a)-[:LIVES_IN]->(b:Country) WHERE b.population > 10000000 RETURN a, b"
    "CREATE (n:Person {name: 'Charlie', age: 35})"
    "MATCH (n)-[:LIKES]->(m:Food) WHERE m.type = 'Pizza' RETURN n, m"
    "MATCH (a)-[:FRIEND]->(b) RETURN a, b"
    "MATCH (n)-[:ATTENDED]->(m:Event) WHERE m.date > '2022-01-01' RETURN n, m"
    "CREATE (n:Person {name: 'David', age: 25})"
    "MATCH (a)-[:OWNS]->(b:House) WHERE a.name = 'Eva' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Pet) WHERE m.species = 'Dog' RETURN n, m"
    "MATCH (a)-[:WORKS_AT]->(b:Company) WHERE b.name = 'TechCorp' RETURN a, b"
    "CREATE (n:Person {name: 'Grace', age: 28, gender: 'Female'})"
    "MATCH (n)-[:FOLLOWS]->(m:Person) WHERE m.name = 'John' RETURN n, m"
    "MATCH (a)-[:LIKES]->(b:Movie) WHERE b.genre = 'Action' RETURN a, b"
    "CREATE (n:Book {title: 'The Art of Programming', author: 'Donald Knuth'})"
    "MATCH (a:Person)-[:LIVES_IN]->(b:City) WHERE b.name = 'New York' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Skill) WHERE m.type = 'Programming' RETURN n, m"
    "CREATE (n:Person {name: 'Sophie', age: 32, occupation: 'Engineer'})"
    "MATCH (a)-[:VISITED]->(b:Country) WHERE b.name = 'Japan' RETURN a, b"
    "MATCH (n)-[:ATTENDED]->(m:Event) WHERE m.type = 'Conference' RETURN n, m"
    "MATCH (a)-[:OWNS]->(b:Car) WHERE b.brand = 'Tesla' RETURN a, b"
    "CREATE (n:Person {name: 'Frank', age: 40, hobby: 'Photography'})"
    "MATCH (n)-[:FRIEND]->(m:Person) WHERE m.name = 'Ella' RETURN n, m"
    "MATCH (a)-[:LIKES]->(b:Food) WHERE b.type = 'Sushi' RETURN a, b"
    "CREATE (n:Person {name: 'Isaac', age: 27, field: 'Data Science'})"
    "MATCH (a)-[:WORKS_AT]->(b:Company) WHERE b.name = 'InnovateTech' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Pet) WHERE m.species = 'Cat' RETURN n, m"
    "CREATE (n:Movie {title: 'The Matrix', year: 1999, genre: 'Sci-Fi'})"
    "MATCH (a)-[:LIVES_IN]->(b:City) WHERE b.name = 'San Francisco' RETURN a, b"
    "MATCH (n)-[:FOLLOWS]->(m:Person) WHERE m.name = 'Alice' RETURN n, m"
    "MATCH (a)-[:VISITED]->(b:City) WHERE b.name = 'Berlin' RETURN a, b"
    "MATCH (n)-[:FRIEND]->(m:Person) WHERE m.name = 'Charlie' RETURN n, m"
    "CREATE (n:Person {name: 'Olivia', age: 35, hobby: 'Cooking'})"
    "MATCH (a)-[:LIKES]->(b:Movie) WHERE b.title = 'Inception' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Skill) WHERE m.type = 'Design' RETURN n, m"
    "MATCH (a)-[:WORKS_AT]->(b:Company) WHERE b.name = 'TechInnovate' RETURN a, b"
    "CREATE (n:Book {title: 'The Power of Now', author: 'Eckhart Tolle'})"
    "MATCH (a:Person)-[:LIVES_IN]->(b:City) WHERE b.name = 'Tokyo' RETURN a, b"
    "MATCH (n)-[:VISITED]->(m:Country) WHERE m.name = 'Australia' RETURN n, m"
    "CREATE (n:Person {name: 'Samuel', age: 28, occupation: 'Artist'})"
    "MATCH (a)-[:ATTENDED]->(b:Event) WHERE b.type = 'Workshop' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Pet) WHERE m.species = 'Parrot' RETURN n, m"
    "CREATE (n:Movie {title: 'The Shawshank Redemption', year: 1994, genre: 'Drama'})"
    "MATCH (a)-[:LIKES]->(b:Food) WHERE b.type = 'Italian' RETURN a, b"
    "MATCH (n)-[:FRIEND]->(m:Person) WHERE m.name = 'Grace' RETURN n, m"
    "CREATE (n:Person {name: 'Liam', age: 32, field: 'Software Development'})"
    "MATCH (a)-[:WORKS_AT]->(b:Company) WHERE b.name = 'DataTech' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Pet) WHERE m.species = 'Rabbit' RETURN n, m"
    "CREATE (n:Book {title: 'Sapiens', author: 'Yuval Noah Harari'})"
    "MATCH (a:Person)-[:LIVES_IN]->(b:City) WHERE b.name = 'Los Angeles' RETURN a, b"
    "MATCH (n)-[:VISITED]->(m:Country) WHERE m.name = 'Canada' RETURN n, m"
    "CREATE (n:Person {name: 'Sophia', age: 26, hobby: 'Gardening'})"
    "MATCH (a)-[:LIKES]->(b:Movie) WHERE b.title = 'The Dark Knight' RETURN a, b"
    "MATCH (n)-[:FRIEND]->(m:Person) WHERE m.name = 'David' RETURN n, m"
    "CREATE (n:Book {title: 'Thinking, Fast and Slow', author: 'Daniel Kahneman'})"
    "MATCH (a)-[:ATTENDED]->(b:Event) WHERE b.type = 'Conference' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Skill) WHERE m.type = 'Machine Learning' RETURN n, m"
    "CREATE (n:Person {name: 'Emily', age: 30, occupation: 'Data Analyst'})"
    "MATCH (a)-[:WORKS_AT]->(b:Company) WHERE b.name = 'InnoSoft' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Pet) WHERE m.species = 'Dog' RETURN n, m"
    "CREATE (n:Movie {title: 'Pulp Fiction', year: 1994, genre: 'Crime'})"
    "MATCH (a:Person)-[:LIVES_IN]->(b:City) WHERE b.name = 'Barcelona' RETURN a, b"
    "MATCH (n)-[:VISITED]->(m:Country) WHERE m.name = 'Brazil' RETURN n, m"
    "CREATE (n:Person {name: 'Lucas', age: 33, hobby: 'Traveling'})"
    "MATCH (a)-[:LIKES]->(b:Food) WHERE b.type = 'Mexican' RETURN a, b"
    "CREATE (n:Person {name: 'Sophie', age: 28, hobby: 'Reading'})"
    "MATCH (a)-[:FOLLOWS]->(b:Person) WHERE b.name = 'Eva' RETURN a, b"
    "MATCH (n)-[:VISITED]->(m:City) WHERE m.name = 'Sydney' RETURN n, m"
    "CREATE (n:Book {title: 'The Hitchhiker''s Guide to the Galaxy', author: 'Douglas Adams'})"
    "MATCH (a)-[:HAS]->(b:Skill) WHERE b.type = 'Artificial Intelligence' RETURN a, b"
    "MATCH (n)-[:LIKES]->(m:Movie) WHERE m.title = 'The Godfather' RETURN n, m"
    "CREATE (n:Person {name: 'Alice', age: 30, occupation: 'Software Engineer'})"
    "MATCH (a)-[:WORKS_AT]->(b:Company) WHERE b.name = 'InnovateNow' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Pet) WHERE m.species = 'Fish' RETURN n, m"
    "CREATE (n:Movie {title: 'Forrest Gump', year: 1994, genre: 'Drama'})"
    "MATCH (a)-[:VISITED]->(b:Country) WHERE b.name = 'Germany' RETURN a, b"
    "MATCH (n)-[:FRIEND]->(m:Person) WHERE m.name = 'Robert' RETURN n, m"
    "CREATE (n:Book {title: '1984', author: 'George Orwell'})"
    "MATCH (a)-[:ATTENDED]->(b:Event) WHERE b.type = 'Concert' RETURN a, b"
    "MATCH (n)-[:HAS]->(m:Skill) WHERE m.type = 'Frontend Development' RETURN n, m"
    "CREATE (n:Person {name: 'Mia', age: 29, hobby: 'Photography'})"
    "MATCH (a)-[:LIKES]->(b:Movie) WHERE b.title = 'The Shawshank Redemption' RETURN a, b"
    "MATCH (n)-[:FOLLOWS]->(m:Person) WHERE m.name = 'Lily' RETURN n, m"
    "CREATE (n:Book {title: 'To Kill a Mockingbird', author: 'Harper Lee'})"
)

# Function to run benchmark for a single query
run_benchmark() {
    query="$1"
    total_time=0

    for ((i = 1; i <= 100; i++)); do
        # Measure the time taken for each query execution
        query_time=$( { time -p curl -s -X POST -H "Content-Type: application/json" -d '{"query":"'"$query"'"}' http://localhost:8081; } 2>&1 | grep real | awk '{print $2}' )
        total_time=$(echo "$total_time + $query_time" | bc)
    done

    # Calculate the average time
    average_time=$(echo "scale=4; $total_time / 100" | bc)
    echo "| $query | $average_time sec | $total_time sec"
}

# Display table header
echo "+-----------------------------------------------+"
echo "| Cypher Query | Average Time   | Total Time"
echo "+-----------------------------------------------+"

# Run benchmark for each query
for query in "${cypher_queries[@]}"; do
    run_benchmark "$query"
done

# Display table footer
echo "+-----------------------------------------------+"