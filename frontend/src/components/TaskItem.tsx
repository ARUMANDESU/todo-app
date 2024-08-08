import React from 'react';
import {domain} from "../../wailsjs/go/models";
import {Card, CardContent} from "@/components/ui/card";
import {Badge} from "@/components/ui/badge";
import {Button} from "@/components/ui/button";
import {CheckIcon, Delete, DeleteIcon, Trash} from "lucide-react";
import TaskPriority = domain.TaskPriority;
import {Input} from "@/components/ui/input";
import {DeleteTask} from "../../wailsjs/go/main/App";
import {toast} from "sonner";

export type TaskItemProps = {
    task: domain.Task;
    onClick: (task: domain.Task) => void;
    isCurrent: boolean;
}

const priorityColors: { [key: string]: string } = {
    high: "text-red-500",
    medium: "text-yellow-500",
    low: "text-green-500",
    normal: "text-blue-500"
};

function TaskItem({task, onClick, isCurrent}: TaskItemProps) {

    const markTaskDone = (task: domain.Task) => {

    }

    const deleteTask = (task: domain.Task) => {
        DeleteTask(task.id)
            .then(() => {
                toast.success("Task deleted successfully")
            })
            .catch((err) => {
                toast.error("Failed to delete task")
            })
    }

    return (
        <>
            <Card
                onClick={() => onClick(task)}
                className={`cursor-pointer ${ isCurrent ? "bg-muted" : "hover:bg-muted"}`}
            >
                <CardContent className="flex justify-between items-center">
                    <div className="flex flex-row justify-start">
                        <Input type="checkbox"/>
                        <div className="font-medium">{task.title}</div>
                        <div className="text-sm text-muted-foreground">
                            Due: {task.due_date}
                            <span className="ml-2">
                                {priorityBadge(task.priority)}
                            </span>
                        </div>
                    </div>
                    <Button>
                        <Trash />
                    </Button>
                </CardContent>
            </Card>
        </>
    );
}

function priorityBadge(priority: TaskPriority) {
    switch (priority) {
        case TaskPriority.NONE:
            return (
                <Badge variant="default" className="bg-muted text-muted-foreground">
                    None
                </Badge>
            );
        case TaskPriority.LOW:
            return (
                <Badge variant="default" className="bg-green-500 text-green-50">
                    Low
                </Badge>
            );
        case TaskPriority.MEDIUM:
            return (
                <Badge variant="default" className="bg-yellow-500 text-yellow-50">
                    Medium
                </Badge>
            );
        case TaskPriority.HIGH:
            return (
                <Badge variant="default" className="bg-red-500 text-red-50">
                    High
                </Badge>
            );
    }
}

export default TaskItem;