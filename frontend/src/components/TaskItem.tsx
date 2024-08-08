import React from 'react';
import {domain} from "../../wailsjs/go/models";
import {Card, CardContent} from '@/components/ui/card';
import {Square, SquareX, Trash} from "lucide-react";
import {DeleteTask, UpdateTask} from "../../wailsjs/go/main/App";
import {toast} from "sonner";
import {PriorityBadge} from "@/components/PriorityBadge";
import TaskStatus = domain.TaskStatus;

export type TaskItemProps = {
    task: domain.Task;
    onClick: (task: domain.Task) => void;
    onUpdate: (task: domain.Task) => void;
    onDelete: (task: domain.Task) => void;
    isCurrent: boolean;
}

const priorityColors: { [key: string]: string } = {
    high: "text-red-500 hover:bg-red-100",
    medium: "text-yellow-500 hover:bg-yellow-100",
    low: "text-green-500 hover:bg-green-100",
    normal: "text-blue-500 hover:bg-blue-100",
};

function TaskItem({task, onClick, isCurrent, onUpdate, onDelete}: TaskItemProps) {

    const markTaskDone = (task: domain.Task) => {
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

    const deleteTask = (task: domain.Task) => {
        DeleteTask(task.id)
            .then(() => {
                onDelete(task)
                toast.success("Task deleted successfully")
            })
            .catch((err) => {
                if (err.message !== "cancelled") {
                    return
                }
                toast.error(`Failed to delete task: ${err}`)
            })
    }

    return (
        <>
            <Card className={`cursor-pointer ${ isCurrent ? "bg-muted" : "hover:bg-muted"}`}>
                <CardContent className="flex justify-between items-center p-3">
                    <div className="flex justify-start w-full" onClick={() => onClick(task)}>
                        {task.status === TaskStatus.DONE ?
                            (<SquareX  className="text-muted-foreground hover:bg-primary/20" onClick={() => markTaskDone(task)}/> ) :
                            (<Square className={priorityColors[task.priority]} onClick={() => markTaskDone(task)}/>)
                        }
                        <div className="font-medium">{task.title}</div>
                        <div className="text-sm text-muted-foreground">
                            Due: {task.due_date}
                            <span className="ml-2">
                                <PriorityBadge priority={task.priority}/>
                            </span>
                        </div>
                    </div>
                    <Trash  onClick={()=>{deleteTask(task)}}/>
                </CardContent>
            </Card>
        </>
    );
}

export default TaskItem;