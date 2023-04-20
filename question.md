1. Explain what RESTful API is 

REST stands for Representational State Transfer. It is a style of web architecture that for building web services. RESTful APIs should be simple, lightweight, and scalable, which use HTTP requests to transfer data and state between clients and servers.

The core concept of a RESTful API is resources. Each resource can have multiple representations, such as XML, JSON, or HTML. RESTful APIs define a set of standard HTTP methods, such as GET, POST, PUT, and DELETE, which are used to perform operations on resources.

The key characteristics of a RESTful API include:

Client-server architecture: The client and server are separated by a uniform interface, which allows them to evolve independently.

Stateless: Each request from the client to the server must contain all the information necessary to understand the request, and the server must not store any client context between requests.

Cacheable: Responses from the server must be explicitly marked as cacheable or non-cacheable.

Layered system: The architecture can be composed of multiple layers, such as load balancers or caching servers, without affecting the API's behavior.

Uniform interface: The API should have a consistent, predictable interface that uses HTTP methods and URIs to represent resources and operations on those resources.

Overall, a RESTful API provides a standardized way for clients and servers to communicate over the web. By following the principles of RESTful design, developers can create APIs that are reliable, scalable, and easy to use.

2. How do you form an api endpoint to query with below resources
    Page: 10
    Book: odyssey
    Library: GreatLibrary   

GET /api/library/GreatLibrary/book/odyssey/page/10

GET: The HTTP method used to retrieve data from the server.
/api: The base URL path for the API.
/library/GreatLibrary: The resource path that identifies the library.
/book/odyssey: The resource path that identifies the book.
/page/10: The resource path that identifies the page number.

3. Explain what MVC is 

MVC stands for Model-View-Controller, and it is a popular architectural pattern used in software development to separate an application into three interconnected components:

Model: This component represents the data and the business logic of the application. It handles the data storage, retrieval, and manipulation. It is responsible for defining the rules and behaviors that govern the application's behavior.

View: This component is responsible for rendering the user interface of the application. It presents the data to the user in a way that is visually appealing and easy to understand. It is responsible for receiving user input and passing it to the controller for processing.

Controller: This component acts as an intermediary between the model and the view. It receives input from the view and translates it into actions that the model can understand. It also updates the view with new data from the model.

By separating the application into these three components, the MVC pattern helps to improve the modularity, maintainability, and scalability of the application. Changes to one component do not affect the other components, making it easier to modify and test the application. The MVC pattern is used in a variety of software applications, including web applications, desktop applications, and mobile applications.

4. Explain what OOP is 

OOP stands for Object-Oriented Programming, and it is a programming paradigm that focuses on creating objects that contain data and methods to manipulate that data. In OOP, an object is an instance of a class, which is a blueprint for creating objects.

The main concepts of OOP include:

Encapsulation: to ensure that the object is used correctly and prevents unwanted access to the object's internal state.

Inheritance: to allows one class to inherit properties and methods from another class. This allows developers to reuse code and reduce the amount of code that needs to be written.

Polymorphism: refers to the ability of an object to take on many forms. This allows developers to create a single interface that can be used to interact with different types of objects.

By using OOP, developers can create more modular, reusable, and maintainable code. OOP is widely used in a variety of programming languages, including Java, Python, and C++.