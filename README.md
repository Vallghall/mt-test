## Тестовое задание
### API reference

POST "/add-fact" adds a relation to between two
people in a DB graph
```json
{
    "person_a": {
        "name": "Roman", 
        "surname": "Gusev",
        "age": 21
    },
    "person_b": {
        "name": "Alex",
        "surname": "Johns",
        "age": 22
    },
    "communication": {
        "type": "SENDS_MESSAGE_TO",
        "description": "Roman sends a letter to Alex"
    }
}
```

GET "/relations" -- returns a graph stored in DB
