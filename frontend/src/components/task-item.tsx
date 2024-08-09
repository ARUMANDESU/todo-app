import React from 'react';
import {domain} from "../../wailsjs/go/models";
import {Card, CardContent} from '@/components/ui/card';
import {Square, SquareX, Trash} from "lucide-react";
import {DeleteTask, UpdateTask} from "../../wailsjs/go/main/App";
import {toast} from "sonner";
import {PriorityBadge} from "@/components/priority-badge";
import TaskStatus = domain.TaskStatus;
import TaskStatusToggle from "@/components/task-status-toggle";

export type TaskItemProps = {
    task: domain.Task;
    onClick: (task: domain.Task) => void;
    onUpdate: (task: domain.Task) => void;
    onDelete: (task: domain.Task) => void;
    isCurrent: boolean;
}

export const priorityColors: { [key: string]: string } = {
    high: "text-red-500 hover:bg-red-100 dark:text-red-300 dark:hover:bg-red-300/30",
    medium: "text-yellow-500 hover:bg-yellow-100 dark:text-yellow-300 dark:hover:bg-yellow-300/30",
    low: "text-green-500 hover:bg-green-100 dark:text-green-300 dark:hover:bg-green-300/30",
    none: "text-blue-500 hover:bg-blue-100 dark:text-blue-300 dark:hover:bg-blue-300/30",
};

function TaskItem({task, onClick, isCurrent, onUpdate, onDelete}: TaskItemProps) {
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
            <Card
                className={`cursor-pointer ${ isCurrent ? "bg-muted" : "hover:bg-muted"}`}
                onClick={() => onClick(task)}
            >
                <CardContent className="flex justify-between items-center p-3">
                    <div className="flex justify-start w-full">
                        <TaskStatusToggle task={task} onClick={markTaskDone}/>
                        <div className={`font-medium px-2.5 ${task.status  === TaskStatus.DONE && " line-through "}`}>{task.title}</div>
                        <span className="ml-2 ">
                            <PriorityBadge priority={task.priority}/>
                        </span>
                    </div>
                    <div className="flex justify-end items-center w-full">
                        <p className="text-muted-foreground text-sm px-2.5">
                            {task.due_date ? new Date(task.due_date).toDateString() : ""}
                        </p>
                        <Trash
                            className={`hover:text-red-500`}
                            onClick={() => {deleteTask(task)}}
                        />
                    </div>

                </CardContent>
            </Card>
        </>
    );
}

export default TaskItem;