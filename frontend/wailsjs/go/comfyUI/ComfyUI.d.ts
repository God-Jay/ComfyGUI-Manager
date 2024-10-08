// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {object} from '../models';
import {context} from '../models';

export function CheckIsServerRunning():Promise<boolean>;

export function GitHeadCommit():Promise<object.Commit>;

export function GitLatestCommit():Promise<void>;

export function GitPull():Promise<string>;

export function GitStatus():Promise<Array<object.Commit>>;

export function Shutdown():Promise<void>;

export function StartServer():Promise<boolean>;

export function StartUp(arg1:context.Context):Promise<void>;

export function StopServer():Promise<boolean>;

export function StoreOutput(arg1:string,arg2:string):Promise<void>;
