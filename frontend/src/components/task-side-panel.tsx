import React, {useEffect, useState} from 'react';
import {domain} from "../../wailsjs/go/models";
import {Card, CardContent, CardHeader} from "@/components/ui/card";
import {debounce} from "ts-debounce";
import {UpdateTask} from "../../wailsjs/go/main/App";
import {toast} from "sonner";
import {Input} from "@/components/ui/input";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue
} from "@/components/ui/select";
import {Flag} from "lucide-react";
import TaskPriority = domain.TaskPriority;

export type TaskSidePanelProps = {
    task: domain.Task | null;
    onUpdate: (task: domain.Task) => void;
}

function TaskSidePanel({task, onUpdate}: TaskSidePanelProps) {
    const [title, setTitle] = useState<string>(task ? task.title : "");
    const [titleError, setTitleError] = useState<string>("");

    const [priority, setPriority] = useState<TaskPriority>(task ? task.priority : TaskPriority.NONE);
    const [dueDate, setDueDate] = useState<string>(
        task?.due_date ?
            new Date(task.due_date).toISOString().split('T')[0]
            : Date.now().toString()

);

    useEffect(() => {
        setTitle(task ? task.title : "");
        setPriority(task ? task.priority : TaskPriority.NONE);
        setDueDate(
            task?.due_date ?
                new Date(task.due_date).toISOString().split('T')[0]
                : Date.now().toString()
        );
    }, [task]);


    const onTitleChange = (e: any) => {
        if (!task) {
            toast.error("No task selected");
            return;
        }

        setTitle(e.target.value);
        setTitleError("");

        if (!e.target.value) {
            setTitleError("Title cannot be empty");
            return;
        }
        if (e.target.value.length > 250 || e.target.value.length < 3) {
            setTitleError("Title must be between 3 and 250 characters");
            return;
        }

        debounce( async () => {
            try {
                const updatedTask = await UpdateTask(
                    domain.UpdateTaskRequest.createFrom(
                        {id:task.id, title: e.target.value}
                    )
                );
                onUpdate(updatedTask);
            } catch (e) {
                toast.error("Failed to update task's title: " + e);
                console.error(e);
            }

        }, 500)();
    }


    const onPriorityChange = (value: TaskPriority) => {
        if (!task) {
            toast.error("No task selected");
            return;
        }
        setPriority(value);

        debounce(async () => {
            try {
                const updatedTask = await UpdateTask(
                    domain.UpdateTaskRequest.createFrom(
                        {id: task.id, priority: value}
                    )
                );
                onUpdate(updatedTask);
            } catch (e) {
                toast.error("Failed to change priority: " + e);
                console.error(e);
            }
        }, 230)();
    }

    const onDueDateChange = (e: any) => {
        if (!task) {
            toast.error("No task selected");
            return;
        }

        setDueDate(e.target.value);

        debounce(async () => {
            try {
                const updatedTask = await UpdateTask(
                    domain.UpdateTaskRequest.createFrom(
                        {
                            id: task.id,
                            due_date: new Date(e.target.value).toISOString()
                        }
                    )
                );
                setDueDate(updatedTask.due_date);
                onUpdate(updatedTask);

            } catch (e) {
                toast.error("Failed to change due date: " + e);
                console.error(e);
            }

        }, 500)();
    }

    return (
        <div>
            {task ? (
                <Card className="sticky top-8">
                    <CardHeader>
                        <p className="text-sm text-red-500">{titleError}</p>
                        <Input
                            type="text"
                            value={title}
                            onChange={onTitleChange}
                            placeholder="Task title"
                            className="text-lg font-bold w-full border-0"
                        />
                    </CardHeader>
                    <CardContent className="space-y-4">
                        <div>
                            <div className="text-sm text-muted-foreground">
                                Priority:
                                <Select value={priority} onValueChange={onPriorityChange}>
                                    <SelectTrigger className="w-full">
                                        <SelectValue placeholder="Select priority"/>
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectGroup>
                                            <SelectLabel>Priority</SelectLabel>
                                            <SelectItem value={TaskPriority.NONE}>
                                                <div className="flex items-center">
                                                    <Flag className="h-4 w-4 mr-2" />
                                                    <span>None</span>
                                                </div>
                                            </SelectItem>
                                            <SelectItem value={TaskPriority.LOW}>
                                                <div className="flex items-center">
                                                    <Flag className="h-4 w-4 mr-2 text-green-500 dark:text-green-300" />
                                                    <span>Low</span>
                                                </div>
                                            </SelectItem>
                                            <SelectItem value={TaskPriority.MEDIUM}>
                                                <div className="flex items-center">
                                                    <Flag className="h-4 w-4 mr-2 text-yellow-500 dark:text-yellow-300" />
                                                    <span>Medium</span>
                                                </div>
                                            </SelectItem>
                                            <SelectItem value={TaskPriority.HIGH}>
                                                <div className="flex items-center">
                                                    <Flag className="h-4 w-4 mr-2 text-red-500 dark:text-red-300" />
                                                    <span>High</span>
                                                </div>
                                            </SelectItem>
                                        </SelectGroup>
                                    </SelectContent>
                                </Select>
                            </div>
                            <div className={`text-sm ${task.status === "done" ? "placeholder:text-muted-foreground" : dueDate < Date.now().toString() && task.due_date ? "text-red-500" : "text-muted-foreground"}`}>
                                Due Date:
                                <Input
                                    type="date"
                                    value={dueDate}
                                    onChange={onDueDateChange}
                                    className="w-full"
                                />
                            </div>


                        </div>
                    </CardContent>
                </Card>
            ) : (
                <Card className="sticky top-8">
                    <CardContent className="p-4">
                        <p className="text-muted-foreground">Select a task to view its details.</p>
                    </CardContent>
                </Card>
            )}
        </div>
    );
}

export default TaskSidePanel