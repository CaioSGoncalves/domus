export type Task = {
  id: number;
  title: string;
};

export type Column = {
  id: number;
  title: string;
  tasks: Task[];
};

export type Panel = {
  id: number;
  name: string;
  columns: Column[];
};

export async function getPanels(): Promise<Panel[]> {
  const res = await fetch("/api/tasks");
  return res.json();
}

export async function moveTask(
  panelId: number,
  columnId: number,
  id: number,
): Promise<void> {
  await fetch("/api/tasks/move", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id, columnId }),
  });
}

export async function addTask(
  panelId: number,
  columnId: number,
  title: string,
): Promise<Task | null> {
  if (!title) return null;

  const response = await fetch("/api/tasks", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ panelId, columnId, title }),
  });

  return (await response.json()) as Task;
}
