tasks = []

def add_task():
    task = input("Enter a task: ")
    tasks.append(task)
    print("Task added!")

def view_tasks():
    if not tasks:
        print("No tasks yet.")
        return

    print("\nYour Tasks:")
    for i, task in enumerate(tasks, 1):
        print(f"{i}. {task}")

def delete_task():
    view_tasks()
    try:
        num = int(input("Enter task number to delete: "))
        removed = tasks.pop(num - 1)
        print("Removed:", removed)
    except:
        print("Invalid selection")

while True:

    print("\nTask Tracker")
    print("1. Add Task")
    print("2. View Tasks")
    print("3. Delete Task")
    print("4. Exit")

    choice = input("Choose: ")

    if choice == "1":
        add_task()

    elif choice == "2":
        view_tasks()

    elif choice == "3":
        delete_task()

    elif choice == "4":
        print("Goodbye!")
        break

    else:
        print("Invalid option")