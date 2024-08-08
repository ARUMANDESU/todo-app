// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {domain} from '../models';

export function CreateTask(arg1:domain.CreateTaskRequest):Promise<domain.Task>;

export function DeleteTask(arg1:number):Promise<void>;

export function GetAllTasks():Promise<Array<domain.Task>>;

export function GetTaskByID(arg1:number):Promise<domain.Task>;

export function Greet(arg1:string):Promise<string>;

export function UpdateTask(arg1:domain.UpdateTaskRequest):Promise<domain.Task>;
