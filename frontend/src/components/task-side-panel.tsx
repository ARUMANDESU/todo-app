import React, {useEffect, useState} from 'react';
import {domain} from "../../wailsjs/go/models";
import {Card, CardContent, CardHeader} from "@/components/ui/card";
import {debounce} from "ts-debounce";
import {UpdateTask} from "../../wailsjs/go/main/App";
import {toast} from "sonner";
import {Input} from "@/components/ui/input";
import PrioritySelect from "@/components/priority-select";
import TaskPriority = domain.TaskPriority;
import TaskStatusToggle from "@/components/task-status-toggle";
import TaskStatus = domain.TaskStatus;
import {Separator} from "@/components/ui/separator";
import {Textarea} from "@/components/ui/textarea";
import {Badge} from "@/components/ui/badge";
import {Button} from "@/components/ui/button";
import {XIcon} from "lucide-react";

export type TaskSidePanelProps = {
    task: domain.Task | null;
    onUpdate: (task: domain.Task) => void;
}

function TaskSidePanel({task, onUpdate}: TaskSidePanelProps) {
    if (task === null) {
        return (
            <Card className="sticky top-8">
                <CardContent className="p-4">
                    <p className="text-muted-foreground">Select a task to view its details.</p>
                </CardContent>
            </Card>
        )
    }

    const [title, setTitle] = useState<string>(task.title);
    const [titleError, setTitleError] = useState<string>("");
    const [description, setDescription] = useState<string>(task.description);
    const [descriptionError, setDescriptionError] = useState<string>("");
    const [priority, setPriority] = useState<TaskPriority>(task.priority);
    const [dueDate, setDueDate] = useState<string>(
        task.due_date ?
            new Date(task.due_date).toISOString().split('T')[0]
            : Date.now().toString()
    );
    const [tags, setTags] = useState<string[]>(task.tags || []);

    const [tagInput, setTagInput] = useState<string>("");

    const onTitleChange = (e: any) => {
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

    const onDescriptionChange = (e: any) => {
        setDescription(e.target.value);
        setDescriptionError("");

        if (e.target.value.length > 1000) {
            setDescriptionError("Description must be less than 1000 characters");
            return;
        }

        debounce(async () => {
            try {
                const updatedTask = await UpdateTask(
                    domain.UpdateTaskRequest.createFrom(
                        {id: task.id, description: e.target.value, tags: tags}
                    )
                );
                onUpdate(updatedTask);
            } catch (e) {
                toast.error("Failed to update task's description: " + e);
                console.error(e);
            }

        }, 500)();
    }

    const onRemoveTag = async (tag: string) => {
        try {
            const updatedTask = await UpdateTask(
                domain.UpdateTaskRequest.createFrom(
                    {id: task.id, description: description, tags: tags.filter((t) => t !== tag)}
                )
            );
            onUpdate(updatedTask);
        } catch (e) {
            toast.error("Failed to update task's description: " + e);
            console.error(e);
        }
    }

    const addTag = async (tag: string) => {
        if (tags.includes(tag)) {
            console.log("Tag already exists");
            return;
        }
        try {
            const updatedTask = await UpdateTask(
                domain.UpdateTaskRequest.createFrom(
                    {id: task.id, description: description, tags: [...tags, tag]}
                )
            );
            onUpdate(updatedTask);
        } catch (e) {
            toast.error("Failed to add new tag: " + e);
            console.error(e);
        }
    }

    const markTaskDone = () => {
        const request = domain.UpdateTaskRequest.createFrom(
            {
                id: task.id,
                status: task.status === TaskStatus.DONE ? TaskStatus.TODO : TaskStatus.DONE
            })
        UpdateTask(request).then(onUpdate).catch((err) => {
            toast.error(`Failed to update task: ${err}`)
            console.log(err)
        })
    }

    useEffect(() => {
        setTitle(task.title);
        setPriority(task.priority);
        setDueDate(
            task.due_date ?
                new Date(task.due_date).toISOString().split('T')[0]
                : Date.now().toString()
        );
        setDescription(task.description);
        setTags(task.tags || []);
        setTitleError("")
        setDescriptionError("")
    }, [task]);

    return (
        <div>
            <Card className="sticky top-8 min-h-[50vh] max-h-full">
                <CardHeader className="flex flex-row p-1 px-2 justify-between" >
                    <div className="flex flex-row justify-start content-center gap-2">
                        <div className="content-center">
                            <TaskStatusToggle task={task} onClick={markTaskDone}/>
                        </div>
                        <Separator orientation="vertical" className="h-[1.1rem] absolute left-8 top-5 w-[2px] ml-0.5"/>
                        <div
                            className={
                                `text-sm content-center py-1                                                                                
                                ${task.status === "done" ? 
                                    "placeholder:text-muted-foreground" 
                                    : dueDate < Date.now().toString() && task.due_date ? 
                                        "text-red-500" 
                                        : "text-muted-foreground"
                                }`}
                        >
                            <Input
                                type="date"
                                value={dueDate}
                                onChange={onDueDateChange}
                                className="w-full"
                            />
                        </div>
                    </div>
                    <div className="text-sm text-muted-foreground">
                        <PrioritySelect name="priority" priority={priority} setPriority={onPriorityChange}/>
                    </div>
                </CardHeader>
                <Separator orientation="horizontal"/>
                <CardContent className="space-y-4">
                    <div>
                        <Input
                            type="text"
                            value={title}
                            onChange={onTitleChange}
                            placeholder="Task title"
                            className={`p-1 mt-2 text-lg font-bold w-full border-0 ${titleError ? "border-red-500 text-red-500" : ""}`}
                        />
                        <p className="text-sm text-red-500 pt-1">{titleError}</p>
                    </div>
                    <div>
                        <p className="text-sm text-red-500">{descriptionError}</p>
                        <Textarea
                            value={description}
                            onChange={onDescriptionChange}
                            className="w-full h-32 p-2 border-0"
                            placeholder="Task description"
                        />
                    </div>
                    <div className="h-full content-end">
                        <div className="flex flex-wrap gap-2 mt-2">
                            {task.tags?.map((tag, tagIndex) => (
                                <Badge key={tagIndex} className="bg-muted text-muted-foreground hover:bg-muted hover:text-muted-foreground">
                                    {tag}
                                    <XIcon className="w-4 h-4 hover:text-red-500" onClick={()=>onRemoveTag(tag)}/>
                                </Badge>
                            ))}
                            <Input
                                placeholder="Add tag"
                                value={tagInput}
                                onChange={(e) => setTagInput(e.target.value)}
                                onKeyDown={(e) => {
                                    if (e.key === "Enter") {
                                        addTag(tagInput);
                                        setTagInput("");
                                    }
                                }}
                                className="w-auto h-[1.2rem] bg-muted text-muted-foreground rounded-md px-2 py-1 text-sm"
                            />
                        </div>
                    </div>
                </CardContent>
            </Card>
        </div>
    );
}

export default TaskSidePanel