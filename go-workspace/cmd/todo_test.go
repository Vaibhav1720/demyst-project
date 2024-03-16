package todo

import (
    "testing"
    "os"
    "io/ioutil"
    "encoding/json"
    "reflect"
)

func TestAdd(t *testing.T) {
    todos := Todos{}
    todos.Add("task1")

    if len(todos) != 1 {
        t.Errorf("Expected 1 item, got %d", len(todos))
    }

    if todos[0].Title != "task1" {
        t.Errorf("Expected title to be 'task1', got %s", todos[0].Title)
    }
}

func TestComplete(t *testing.T) {
    todos := Todos{{Title: "task1", Completed: false}}
    err := todos.Complete(1)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if !todos[0].Completed {
        t.Errorf("Expected task to be completed")
    }
}

func TestDelete(t *testing.T) {
    todos := Todos{{Title: "task1"}, {Title: "task2"}}
    err := todos.Delete(1)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if len(todos) != 1 {
        t.Errorf("Expected 1 item, got %d", len(todos))
    }

    if todos[0].Title != "task2" {
        t.Errorf("Expected title to be 'task2', got %s", todos[0].Title)
    }
}

func TestLoadAndStore(t *testing.T) {
    // Create temporary file
    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name())

    // Write data to temporary file
    data := Todos{{Title: "task1", Completed: false}}
    err = json.NewEncoder(tmpfile).Encode(data)
    if err != nil {
        t.Fatal(err)
    }

    // Close the file
    if err := tmpfile.Close(); err != nil {
        t.Fatal(err)
    }

    // Load from the temporary file
    todos := Todos{}
    err = todos.Load(tmpfile.Name())
    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if !reflect.DeepEqual(data, todos) {
        t.Errorf("Loaded data does not match stored data")
    }

    // Store to a new file
    newfile := tmpfile.Name() + ".new"
    err = todos.Store(newfile)
    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    // Load from the new file and compare
    newTodos := Todos{}
    err = newTodos.Load(newfile)
    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if !reflect.DeepEqual(data, newTodos) {
        t.Errorf("Loaded data does not match stored data")
    }
}

func TestPrintFirst20EvenNumberedTodos(t *testing.T) {
    todos := Todos{}
    for i := 1; i <= 40; i++ {
        todos = append(todos, item{Title: fmt.Sprintf("task%d", i)})
    }

    // Redirect stdout for testing
    stdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    todos.PrintFirst20EnenNumberedTodos()
    w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = stdout

    expected := ""
    for i := 2; i <= 40; i += 2 {
        expected += fmt.Sprintf("%d - task%d -> Pending \n", i, i)
    }

    if string(out) != expected {
        t.Errorf("Output does not match expected")
    }
}

func TestAddMultiple(t *testing.T) {
    todos := Todos{}
    todos.Add("task1")
    todos.Add("task2")

    if len(todos) != 2 {
        t.Errorf("Expected 2 items, got %d", len(todos))
    }

    if todos[1].Title != "task2" {
        t.Errorf("Expected title to be 'task2', got %s", todos[1].Title)
    }
}

func TestCompleteInvalidIndex(t *testing.T) {
    todos := Todos{}
    err := todos.Complete(1)

    if err == nil {
        t.Errorf("Expected an error for invalid index")
    }

    if len(todos) != 0 {
        t.Errorf("Expected no items, got %d", len(todos))
    }
}

func TestDeleteInvalidIndex(t *testing.T) {
    todos := Todos{}
    err := todos.Delete(1)

    if err == nil {
        t.Errorf("Expected an error for invalid index")
    }

    if len(todos) != 0 {
        t.Errorf("Expected no items, got %d", len(todos))
    }
}

func TestLoadInvalidFile(t *testing.T) {
    todos := Todos{}
    err := todos.Load("nonexistent.json")

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if len(todos) != 0 {
        t.Errorf("Expected no items, got %d", len(todos))
    }
}

func TestStoreInvalidFile(t *testing.T) {
    todos := Todos{{Title: "task1"}}
    err := todos.Store("/invalid/path/todos.json")

    if err == nil {
        t.Errorf("Expected an error for invalid file path")
    }
}

func TestPrintFirst20EvenNumberedTodosLessThan20(t *testing.T) {
    todos := Todos{}
    for i := 1; i <= 10; i++ {
        todos = append(todos, item{Title: fmt.Sprintf("task%d", i)})
    }

    // Redirect stdout for testing
    stdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    todos.PrintFirst20EnenNumberedTodos()
    w.Close()
    out, _ := ioutil.ReadAll(r)
    os.Stdout = stdout

    expected := ""
    for i := 2; i <= 10; i += 2 {
        expected += fmt.Sprintf("%d - task%d -> Pending \n", i, i)
    }

    if string(out) != expected {
        t.Errorf("Output does not match expected")
    }
}

