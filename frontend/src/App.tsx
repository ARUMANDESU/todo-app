import {useEffect, useState} from 'react';
import {domain} from "../wailsjs/go/models";
import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card";
import {GetAllTasks} from "../wailsjs/go/main/App";
import TaskInput from "@/components/TaskInput";
import TaskItem from "@/components/TaskItem";
import {PriorityBadge} from "@/components/PriorityBadge";
import {Button} from "@/components/ui/button";
import {RefreshCwIcon} from "lucide-react";

function App() {
    const [tasks, setTasks] = useState<domain.Task[]>([])
    const [currentTask, setCurrentTask] = useState<domain.Task | null>(null)

    const handleTaskClick = (task: domain.Task) => {
        setCurrentTask(task)
    }

    const handleRefresh = async () => {
        const tasks = await GetAllTasks()
        setTasks(tasks)
    }

    const handleTaskUpdate = (task: domain.Task) => {
        setTasks(tasks.map((t) => t.id === task.id ? task : t))
    }


    useEffect(() => {
        (async () => {
            const tasks = await GetAllTasks()
            setTasks(tasks)
        })()
    }, [])
    return (
        <div id="App" className="h-screen">
            <div className="grid grid-cols-[1fr_300px] gap-8 p-8">
                <div>
                    <div className="flex items-center justify-between mb-4">
                        <h1 className="text-2xl font-bold">Todo App</h1>
                        <Button variant="outline" onClick={handleRefresh}>
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
                                    .map((task, index) => (
                                        <TaskItem
                                            key={index}
                                            task={task}
                                            onClick={handleTaskClick}
                                            isCurrent={task === currentTask}
                                            onUpdate={handleTaskUpdate}
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
                                        />
                                    ))}
                            </div>
                        </div>
                    </div>
                </div>
                <div>
                    {currentTask ? (
                        <Card className="sticky top-8">
                            <CardHeader>
                                <CardTitle>Current Task</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-4">
                                <div>
                                    <div className="font-medium">{currentTask.title}</div>
                                    <div className="text-sm text-muted-foreground">
                                        Priority:
                                        <PriorityBadge priority={currentTask.priority}/>
                                    </div>
                                    <div className="text-sm text-muted-foreground">Due: {currentTask.due_date}</div>
                                </div>
                            </CardContent>
                        </Card>
                    ) : (
                        <Card className="sticky top-8 text-center">
                            <CardContent>
                                <p className="text-muted-foreground">Select a task to view its details.</p>
                            </CardContent>
                        </Card>
                    )}
                </div>
            </div>
        </div>

    )
}

export default App
