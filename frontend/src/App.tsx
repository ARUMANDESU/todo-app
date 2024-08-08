import {useState} from 'react';
import './App.css';
import {CreateTask, Greet} from "../wailsjs/go/main/App";
import {domain} from "../wailsjs/go/models";

function App() {
    const [tasks, setTasks] = useState<domain.Task[]>([]);
    const [title, setTitle] = useState("")
    const updateTitle = (event: any) => setTitle(event.target.value);
    const createTask = () => {
        const task = domain.CreateTaskRequest.createFrom(
            {
                title: title,
                priority: "normal",
                due_date: null
            }
        );

        CreateTask(task).then(addNewTask);
    }

    const addNewTask = (task: domain.Task) => {
        setTasks([...tasks, task]);
    }

    return (
        <div id="App">
            <div>
                <h1>Todo App</h1>
                <input type="text" placeholder="Enter Task title" value={title} onChange={updateTitle}/>
                <button onClick={createTask}>Create Task</button>
            </div>
            {tasks.map((task, index) => (
                <div key={index} className="flex flex-col border border-gray-200 p-4 m-4">
                    <h3>{task.title}</h3>
                    <p>{task.status}</p>
                    <p>{task.priority}</p>
                    {task.due_date && <p>{task.due_date}</p>}
                </div>
            ))}
        </div>
    )
}

export default App
