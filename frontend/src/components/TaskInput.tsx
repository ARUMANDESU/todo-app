import React, {useState} from 'react';
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
import {Badge} from "@/components/ui/badge";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {Button} from "@/components/ui/button";
import {Calendar} from "@/components/ui/calendar";
import {CalendarDaysIcon} from "lucide-react";

export type TaskInputProps = {
    onTaskCreate: (task: domain.Task) => void;
}

function TaskInput({onTaskCreate}: TaskInputProps) {
    const handleSubmit = async (e: any) => {
        e.preventDefault()
        const formData = new FormData(e.target)

        const dueDateString = formData.get("dueDate") as string
        const dueDate = new Date(dueDateString).toISOString()

        const newTask = await CreateTask(domain.CreateTaskRequest.createFrom({
            title: formData.get("title") as string,
            priority: formData.get("priority") as domain.TaskPriority,
            dueDate: dueDate
        }))
        onTaskCreate(newTask)
        e.target.reset()
    }

    return (
        <>
            <form onSubmit={handleSubmit} className="bg-card p-6 rounded-lg shadow-md mb-8 flex items-center gap-4">
                <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                        <Label htmlFor="title">Task Name</Label>
                        <Input id="title" name="title" placeholder="Enter task title" required/>
                    </div>
                    <div className="space-y-2">
                        <Label htmlFor="priority">Priority</Label>
                        {/*select id*/}
                        <Select name="priority" defaultValue="none">
                            <SelectTrigger className="w-full">
                                <SelectValue placeholder="Select priority"/>
                            </SelectTrigger>
                            <SelectContent>
                                <SelectGroup>
                                    <SelectLabel>Priority</SelectLabel>
                                    <SelectItem value="none">
                                        <div className="flex items-center justify-between">
                                            <Badge variant="default" className="bg-muted text-muted-foreground">
                                                None
                                            </Badge>
                                        </div>
                                    </SelectItem>
                                    <SelectItem value="low">
                                        <div className="flex items-center justify-between">
                                            <Badge variant="default" className="bg-green-500 text-green-50">
                                                Low
                                            </Badge>
                                        </div>
                                    </SelectItem>
                                    <SelectItem value="medium">
                                        <div className="flex items-center justify-between">
                                            <Badge variant="default" className="bg-yellow-500 text-yellow-50">
                                                Medium
                                            </Badge>
                                        </div>
                                    </SelectItem>
                                    <SelectItem value="high">
                                        <div className="flex items-center justify-between">
                                            <Badge variant="default" className="bg-red-500 text-red-50">
                                                High
                                            </Badge>
                                        </div>
                                    </SelectItem>
                                </SelectGroup>
                            </SelectContent>
                        </Select>
                    </div>
                </div>
                <div className="space-y-2 mt-4">
                    <Popover>
                        <PopoverTrigger asChild>
                            <Button variant="outline" size="icon">
                                <CalendarDaysIcon/>
                                <span className="sr-only">Due Date</span>
                            </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-auto p-0" align="start">
                            <Calendar mode="single" />
                        </PopoverContent>
                    </Popover>
                </div>
                <div className="mt-6">
                    <Button type="submit" className="w-full">
                        Add Task
                    </Button>
                </div>
            </form>
        </>
    );
}

export default TaskInput;