package main

func main() {

	todos := Todos{}
	store := NewStorage[Todos]("todos.json")
	store.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	store.Save(todos)
}
