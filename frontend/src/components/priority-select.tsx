import React from 'react';
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue
} from "@/components/ui/select";
import {Flag} from "lucide-react";
import {domain} from "../../wailsjs/go/models";
import TaskPriority = domain.TaskPriority;

export type PrioritySelectProps = {
    name: string;
    priority: string;
    setPriority: (priority: TaskPriority) => void;
} & React.ComponentPropsWithoutRef<typeof Select>;

function PrioritySelect({name, priority, setPriority, ...props}: PrioritySelectProps) {
    return (
        <Select name={name} value={priority} onValueChange={setPriority} {...props}>
            <SelectTrigger className="w-full">
                <SelectValue placeholder="Select priority">
                    {priority === "none" && <Flag className="h-4 w-4" />}
                    {priority === "low" && <Flag className="h-4 w-4 text-green-500" />}
                    {priority === "medium" && <Flag className="h-4 w-4 text-yellow-500" />}
                    {priority === "high" && <Flag className="h-4 w-4 text-red-500" />}
                </SelectValue>
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
    )
}

export default PrioritySelect;