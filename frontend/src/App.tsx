import {useEffect, useState} from 'react';
import {domain} from "../wailsjs/go/models";
import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card";
import {GetAllTasks} from "../wailsjs/go/main/App";
import TaskInput from "@/components/TaskInput";
import TaskItem from "@/components/TaskItem";
import {PriorityBadge} from "@/components/PriorityBadge";
import {Button} from "@/components/ui/button";
import {RefreshCwIcon} from "lucide-react";
import TaskSidePanel from "@/components/TaskSidePanel";

const priorityMap: { [key: string]: number } = {
    high: 3,
    medium: 2,
    low: 1,
    none: 0
};

function App() {
    const [tasks, setTasks] = useState<domain.Task[]>([])
    const [currentTask, setCurrentTask] = useState<domain.Task | null>(null)

    const handleTaskClick = (task: domain.Task) => {
        setCurrentTask(task)
    }

    const fetchAllTasks = async () => {
        try {
            const tasks = await GetAllTasks()
            if (tasks === null) {
                return
            }
            setTasks(tasks)
        } catch (e) {
            console.error(e)
        }
    }

    const handleTaskUpdate = (task: domain.Task) => {
        setTasks(tasks.map((t) => t.id === task.id ? task : t))
    }

    const handleTaskDelete = (task: domain.Task) => {
        if (currentTask && currentTask.id === task.id) {
            setCurrentTask(null)
        }
        setTasks(tasks.filter((t) => t.id !== task.id))
    }


    useEffect(() => {
        fetchAllTasks()
    }, [])

    return (
        <div className="min-h-screen h-full">
            <div className="grid grid-cols-[1fr_300px] gap-8 p-8">
                <div>
                    <div className="flex items-center justify-between mb-4">
                        <h1 className="text-2xl font-bold">Todo App</h1>
                        <Button variant="outline" onClick={fetchAllTasks}>
                            <RefreshCwIcon className="h-4 w-4 mr-2"/>
                            Refresh
                        </Button>
                    </div>
                    <TaskInput onTaskCreate={(task) => setTasks([...tasks, task])}/>
                    <div className="grid gap-4">
                        <div>
                            <h2 className="text-lg font-bold mb-2">Todo</h2>
                            <div className="space-y-2">
                                {tasks
                                    .filter((task) => task.status === "todo")
                                    .sort((a, b) => priorityMap[b.priority] - priorityMap[a.priority])
                                    .map((task, index) => (
                                        <TaskItem
                                            key={index}
                                            task={task}
                                            onClick={handleTaskClick}
                                            isCurrent={task === currentTask}
                                            onUpdate={handleTaskUpdate}
                                            onDelete={handleTaskDelete}
                                        />
                                    ))}
                            </div>
                        </div>
                        <div>
                            <h2 className="text-lg font-bold mb-2">Done</h2>
                            <div className="space-y-2">
                                {tasks
                                    .filter((task) => task.status === "done")
                                    .map((task, index) => (
                                        <TaskItem
                                            key={index}
                                            task={task}
                                            onClick={handleTaskClick}
                                            isCurrent={task === currentTask}
                                            onUpdate={handleTaskUpdate}
                                            onDelete={handleTaskDelete}
                                        />
                                    ))}
                            </div>
                        </div>
                    </div>
                </div>
               <TaskSidePanel task={currentTask} onUpdate={handleTaskUpdate}/>
            </div>
        </div>

    )
}

export default App
