export type Task = {
  id: string;
  title: string;
};

export type Column = {
  id: string;
  title: string;
  tasks: Task[];
};

export type Panel = {
  name: string;
  columns: Column[];
};

export async function getPanels(): Promise<Panel[]> {
  const res = await fetch("/api/tasks");
  return res.json();
}

export async function moveTask(id: string, status: string): Promise<void> {
  await fetch("/api/tasks/move", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id, status }),
  });
}
