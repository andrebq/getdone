# Intro

This project is a proof of concept that the ideas behind plumber (github.com/andrebq/plumber) can be used to create real applications.

The idea is that you can write an application as a library and a interface to the user. That way you can reuse all the core concepts of the application without too much burden to carry.

With that in mind, you never write a web application or a GUI application. Instead you write a collection of core concepts (usecases and entities) and expose them via a interface (gui, cli, web, rest, API's, email)

# About getdone

Getdone is a very small and simple Project TODO list. Use it mantain a list of task that needs to be done for a given project.

# Technology

No surprise on the server, it's just GO.

On the client a single HTML file with jQuery and some CSS.

# How it is developed

The most important thing is: write the tests before writing the code. The CreateTask usecase was written and test without a single database call.

This fact shows two important things:

* The logic can be tested very fast.
* ALL CHANGES are tested.
