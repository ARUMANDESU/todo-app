import {useEffect, useState} from 'react';
import {domain} from "../wailsjs/go/models";
import {Label} from "@/components/ui/label";
import {Badge} from "@/components/ui/badge";
import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card";
import {GetAllTasks} from "../wailsjs/go/main/App";
import {Textarea} from "@/components/ui/textarea";
import TaskInput from "@/components/TaskInput";
import TaskItem from "@/components/TaskItem";
import TaskPriority = domain.TaskPriority;

function App() {
    const [tasks, setTasks] = useState<domain.Task[]>([])
    const [currentTask, setCurrentTask] = useState<domain.Task | null>(null)

    const handleTaskClick = (task: domain.Task) => {
        setCurrentTask(task)
    }


    useEffect(() => {
        (async () => {
            const tasks = await GetAllTasks()
            setTasks(tasks)
        })()
    }, [])
    return (
        <div className="h-screen">
            <div className="grid grid-cols-[1fr_300px] gap-8 p-8">
                <div>
                    <h1 className="text-2xl font-bold mb-4">Todo App</h1>
                    <TaskInput onTaskCreate={(task) => setTasks([...tasks, task])}/>
                    <div className="grid gap-4">
                        <div>
                            <h2 className="text-lg font-bold mb-2">Todo</h2>
                            <div className="space-y-2">
                                {tasks
                                    .filter((task) => task.status === "todo")
                                    .map((task, index) => (
                                        <TaskItem task={task} onClick={handleTaskClick}
                                                  isCurrent={task === currentTask}/>
                                    ))}
                            </div>
                        </div>
                        <div>
                            <h2 className="text-lg font-bold mb-2">Done</h2>
                            <div className="space-y-2">
                                {tasks
                                    .filter((task) => task.status === "done")
                                    .map((task, index) => (
                                        <TaskItem task={task} onClick={handleTaskClick}
                                                  isCurrent={task === currentTask}/>
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
                                        {currentTask.priority === TaskPriority.NONE && (
                                            <Badge variant="default" className="bg-muted text-muted-foreground ml-2">
                                                None
                                            </Badge>
                                        )}
                                        {currentTask.priority === TaskPriority.LOW && (
                                            <Badge variant="default" className="bg-green-500 text-green-50 ml-2">
                                                Low
                                            </Badge>
                                        )}
                                        {currentTask.priority === TaskPriority.MEDIUM && (
                                            <Badge variant="default" className="bg-yellow-500 text-yellow-50 ml-2">
                                                Medium
                                            </Badge>
                                        )}
                                        {currentTask.priority === TaskPriority.HIGH && (
                                            <Badge variant="default" className="bg-red-500 text-red-50 ml-2">
                                                High
                                            </Badge>
                                        )}
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
