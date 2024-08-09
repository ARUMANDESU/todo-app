import React, {useRef, useState} from 'react';
import {domain} from "../../wailsjs/go/models";
import {CreateTask} from "../../wailsjs/go/main/App";
import {Label} from "@/components/ui/label";
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
import {Button} from "@/components/ui/button";
import {Flag} from "lucide-react";
import {toast} from "sonner";

export type TaskInputProps = {
    onTaskCreate: (task: domain.Task) => void;
}

function TaskInput({onTaskCreate}: TaskInputProps) {
    const taskInputRef = useRef<HTMLInputElement>(null);
    const [title, setTitle] = useState<string>("");
    const [priority, setPriority] = useState<string>("none");
    const [dueDate, setDueDate] = useState<string>("");

    const handleSubmit = async (e: any) => {
        e.preventDefault();

        if (!title || !priority) {
            toast.error("Please fill all fields");
            return;
        }

        const newTask = await CreateTask(domain.CreateTaskRequest.createFrom({
            title: title,
            priority: priority as domain.TaskPriority,
            due_date: dueDate ? new Date(dueDate).toISOString() : null
        }));
        onTaskCreate(newTask);
        setTitle("");
        setPriority("none");
        setDueDate("");
        taskInputRef.current?.focus();
    }

    return (
        <>
            <form onSubmit={handleSubmit} className="bg-card p-6 rounded-lg shadow-md mb-8 flex items-center gap-4">
                <div className="flex flex-row w-full">
                    <div className="space-y-2 w-full">
                        <Label htmlFor="title">Task Name</Label>
                        <Input
                            id="title"
                            name="title"
                            placeholder="Enter task title"
                            required
                            className="w-full"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                            ref={taskInputRef}
                        />
                    </div>
                    <div className="space-y-2 mx-2">
                        <Label htmlFor="priority">Priority</Label>
                        <Select name="priority" value={priority} onValueChange={setPriority}>
                            <SelectTrigger className="w-full">
                                <SelectValue placeholder="Select priority"/>
                            </SelectTrigger>
                            <SelectContent>
                                <SelectGroup>
                                    <SelectLabel>Priority</SelectLabel>
                                    <SelectItem value="none">
                                        <div className="flex items-center">
                                            <Flag className="h-4 w-4 mr-2" />
                                            <span>None</span>
                                        </div>
                                    </SelectItem>
                                    <SelectItem value="low">
                                        <div className="flex items-center">
                                            <Flag className="h-4 w-4 mr-2 text-green-500" />
                                            <span>Low</span>
                                        </div>
                                    </SelectItem>
                                    <SelectItem value="medium">
                                        <div className="flex items-center">
                                            <Flag className="h-4 w-4 mr-2 text-yellow-500" />
                                            <span>Medium</span>
                                        </div>
                                    </SelectItem>
                                    <SelectItem value="high">
                                        <div className="flex items-center">
                                            <Flag className="h-4 w-4 mr-2 text-red-500" />
                                            <span>High</span>
                                        </div>
                                    </SelectItem>
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                    </div>
                    <div className="space-y-2 mt-8">
                        <Input id="dueDate" name="dueDate" type="date" value={dueDate} onChange={(e) => setDueDate(e.target.value)}/>
                    </div>
                </div>

                <div className="mt-7">
                    <Button type="submit" className="w-full">
                        Add Task
                    </Button>
                </div>
            </form>
        </>
    );
}

export default TaskInput;