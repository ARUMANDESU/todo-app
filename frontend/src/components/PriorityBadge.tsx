import {Badge} from "@/components/ui/badge";
import React from "react";
import {domain} from "../../wailsjs/go/models";
import TaskPriority = domain.TaskPriority;

export function PriorityBadge({priority}: { priority: TaskPriority }) {
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