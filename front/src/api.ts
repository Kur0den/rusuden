import type {Message} from "./type.ts"

export async function listMessages(): Promise<Message[]> {
  const Messages: Message[] = [
    {id: 1, callerName: "test", callerId: "000", created: new Date().toISOString(), url: "/dummy1.wav"},
    {id: 2, callerName: "test", callerId: "001", created: new Date().toISOString(), url: "/dummy1.wav"},
    {id: 3, callerName: "test", callerId: "999000", created: new Date().toISOString(), url: "/dummy1.wav"},
  ]
  return Messages
}
