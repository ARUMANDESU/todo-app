import React, {useEffect, useState} from 'react';
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
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {Button} from "@/components/ui/button";
import {Calendar} from "@/components/ui/calendar";
import {CalendarDaysIcon, Flag} from "lucide-react";
import {toast} from "sonner";

export type TaskInputProps = {
    onTaskCreate: (task: domain.Task) => void;
}

function TaskInput({onTaskCreate}: TaskInputProps) {
    const [dueDate, setDueDate] = useState<string>()
    const [calendarDate, setCalendarDate] = useState<Date | undefined>(new Date())

    const handleSubmit = async (e: any) => {
        e.preventDefault()
        const formData = new FormData(e.target)

        if (!formData.get("title") || !formData.get("priority")) {
            toast.error("Please fill all fields")
            return
        }

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

    useEffect(() => {
        setDueDate(calendarDate?.toISOString())
    }, [calendarDate]);

    return (
        <>
            <form onSubmit={handleSubmit} className="bg-card p-6 rounded-lg shadow-md mb-8 flex items-center gap-4">
                <div className="flex flex-row w-full">
                    <div className="space-y-2 w-full">
                        <Label htmlFor="title">Task Name</Label>
                        <Input id="title" name="title" placeholder="Enter task title" required className="w-full"/>
                    </div>
                    <div className="space-y-2 mx-2">
                        <Label htmlFor="priority">Priority</Label>
                        <Select name="priority" defaultValue="none">
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
                    <div className="space-y-2 mt-6">
                        <Input id="dueDate" name="dueDate" type="date" value={dueDate} className="hidden"/>
                        <Popover>
                            <PopoverTrigger asChild>
                                <Button variant="outline" size="icon">
                                    <CalendarDaysIcon/>
                                    <span className="sr-only">Due Date</span>
                                </Button>
                            </PopoverTrigger>
                            <PopoverContent className="w-auto p-0" align="start">
                                <Calendar
                                    mode="single"
                                    selected={calendarDate}
                                    onSelect={setCalendarDate}
                                />
                            </PopoverContent>
                        </Popover>
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