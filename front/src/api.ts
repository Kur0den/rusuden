import type { Message } from './type.ts'

const BASE = import.meta.env.VITE_API_BASE_URL ?? ''

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
  const res = await fetch(BASE + path, {
    ...init,
    headers: { 'Content-Type': 'application/json', ...init.headers },
    signal: AbortSignal.timeout(10_000),
  })
  if (!res.ok) throw new Error(`${res.status} ${res.statusText}`)
  if (res.status === 204) return undefined as T
  return (await res.json()) as T
}

export const listMessages = () => request<Message[]>('/api/messages')
export const deleteMessage = (id: number) => request<void>(`/messages/${id}`, { method: 'DELETE' })
