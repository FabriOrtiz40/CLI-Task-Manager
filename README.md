# CLI Task Manager (Go)

CLI Task Manager es una herramienta de línea de comandos escrita en **Go** para gestionar tareas (TODOs) desde la terminal.
El proyecto utiliza **Cobra** para la interfaz CLI y **BoltDB (bbolt)** para la persistencia local.

El binario final se llama **CLI-Task-Manager.exe**.

---

## ✨ Funcionalidades

- Agregar tareas
- Listar tareas pendientes
- Marcar tareas como completadas
- Persistencia local automática por usuario
- Funciona desde cualquier directorio

---

## 🛠️ Tecnologías

- Go
- spf13/cobra (CLI)
- go.etcd.io/bbolt (base de datos embebida)

---

## 📦 Instalación

Desde la raíz del proyecto:

```bash
go install .
```

El binario se genera en:

```
C:\Users\<usuario>\go\bin\CLI-Task-Manager.exe
```

Asegurarse de que esa ruta esté incluida en la variable de entorno **PATH**.

---

## ▶️ Uso

### Agregar una tarea

```bash
CLI-Task-Manager add limpiar la cocina
```

Salida:
```
Added "limpiar la cocina" to your task list.
```

---

### Listar tareas pendientes

```bash
CLI-Task-Manager list
```

Salida:
```
You have the following tasks:
1. limpiar la cocina
2. estudiar go
```

---

### Completar una tarea

```bash
CLI-Task-Manager do 1
```

Salida:
```
You have completed task 1.
```

---

## 💾 Persistencia

Las tareas se almacenan automáticamente en un archivo BoltDB ubicado en el directorio HOME del usuario:

- Windows:  
  `C:\Users\<usuario>\.task.db`

La base es única por usuario y no depende del directorio desde el cual se ejecute el comando.

---

## ⚠️ Manejo de errores

- No se utilizan `panic` en la ejecución normal del CLI.
- Los errores se informan con mensajes claros.
- El programa retorna códigos de salida distintos de cero ante fallos.

---

## 📁 Estructura del proyecto

```
CLI-Task-Manager/
├── cmd/
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   └── do.go
├── db/
│   └── db.go
├── main.go
└── go.mod
```

---

## 🎯 Estado del proyecto

- Ejercicio completado
- CLI funcional
- Persistencia validada
- Instalación global operativa

El ejercicio forma parte de Gophercises y sirve para profundizar en el desarrollo de herramientas CLI en Go.