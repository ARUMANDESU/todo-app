import React from 'react';
import {domain} from "../../wailsjs/go/models";
import {Card, CardContent} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import {Trash} from "lucide-react";
import {Input} from "@/components/ui/input";
import {DeleteTask} from "../../wailsjs/go/main/App";
import {toast} from "sonner";
import {PriorityBadge} from "@/components/PriorityBadge";
import TaskStatus = domain.TaskStatus;

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
                toast.error(`Failed to delete task: ${err}`)
            })
    }

    return (
        <>
            <Card className={`cursor-pointer ${ isCurrent ? "bg-muted" : "hover:bg-muted"}`}>
                <CardContent className="flex justify-between items-center">
                    <div className="flex justify-start w-full" onClick={() => onClick(task)}>
                        <Input type="checkbox" checked={task.status == TaskStatus.DONE} onChange={() => markTaskDone(task)}/>
                        <div className="font-medium">{task.title}</div>
                        <div className="text-sm text-muted-foreground">
                            Due: {task.due_date}
                            <span className="ml-2">
                                <PriorityBadge priority={task.priority}/>
                            </span>
                        </div>
                    </div>
                    <Button onClick={()=>{deleteTask(task)}}>
                        <Trash />
                    </Button>
                </CardContent>
            </Card>
        </>
    );
}

export default TaskItem;