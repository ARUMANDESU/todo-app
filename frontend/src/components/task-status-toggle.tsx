import React from 'react';
import {domain} from "../../wailsjs/go/models";
import TaskStatus = domain.TaskStatus;
import {Square, SquareX} from "lucide-react";
import {priorityColors} from "@/components/task-item";

export type TaskStatusToggleProps = {
    task: domain.Task,
    onClick:  React.MouseEventHandler<SVGSVGElement> | undefined
}

function TaskStatusToggle({task, onClick}: TaskStatusToggleProps) {
    if (task.status === TaskStatus.DONE) {
         return (
             <SquareX className="text-muted-foreground hover:bg-primary/20"
                  onClick={onClick}/>
         )
    }

    return (
        <Square className={priorityColors[task.priority]} onClick={onClick}/>
    );
}

export default TaskStatusToggle;