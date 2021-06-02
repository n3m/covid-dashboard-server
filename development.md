# COVID Dashboard

A custom built dashboard that enables the user to view, consume, and filter data related to certain problems in México.
In this build, you'll be able to see different kind of information regarding infections by state, infections by age, and much more...

## Features

- Live-reload Server for Consumers
- Interactive Dashboard

## Tech Stack

**Client:** Javascript

**Server:** Golang, Fiber

## Authors

- [Alan Enrique Maldonado Navarro](https://www.github.com/n3m)
- [Guillermo Gonzalez Mena](https://www.github.com/GG-kun)

## Development Log

- During the first phase of the project we questioned what was going to be our approach, regarding both our tech stack for the tool, and the tool purpose.

- It became clear really quick that for this project we would be using Golang for the server, since it is one of our favourite programming languages and it enables us to use the concurrency applications that it has.

- Regarding the Front-End of our app, we decided to go easy on the design and complexity and went ahead with plain Javascript with D3.Javascript instead of ReactJS and other graphic frameworks.

- Each of us decided to tackle different tasks, but synchronously in order to maintain structure. I (Alan) went ahead and developed the API from Scratch, and made some interesting stuff in order for our project to be fully scalable an easy to use. It featured full dynamic filters for some type of data and ultra-fast performance thanks to the frameworks used and thanks to Golang itself. My partner (Guille) decided to tackle the script that generates the information into a JSON file, and he also went ahead and developed the Front-End Graphics with D3.js that featured live back-end calls for live-results and live-filtering.

- During the development of the Frontend, we met with some difficulties, as it seemed that we forgot a little bit of D3. So in order to fix that, we related back to the practices and projects that we previously made.

- At the end, this project was really beneficial for both of us, since we are going directly into the Data Science field, so we learned a lot of things and tools that will help us along the way.

## API Reference

#### Request Filtered Information

```http
  POST /covid
```

| Body JSON Parameter | Type                                               | Values                                                 |
| :------------------ | :------------------------------------------------- | :----------------------------------------------------- |
| `responseType`      | `string`                                           | ( BYSTATE / BYAGE / BYPRIVPUB / null )                 |
| `sexo`              | `[{"eq": any, "ne": any, "gte": any, "lte": any}]` | Filter Object ("any" can be any value, including null) |
| `edad`              | `[{"eq": any, "ne": any, "gte": any, "lte": any}]` | Filter Object ("any" can be any value, including null) |
| `defunto`           | `[{"eq": any, "ne": any, "gte": any, "lte": any}]` | Filter Object ("any" can be any value, including null) |
| `estadoResidencia`  | `[{"eq": any, "ne": any, "gte": any, "lte": any}]` | Filter Object ("any" can be any value, including null) |
| `origen`            | `[{"eq": any, "ne": any, "gte": any, "lte": any}]` | Filter Object ("any" can be any value, including null) |

## Graph Stories

![Graph 1](https://raw.githubusercontent.com/n3m/covid-dashboard-server/develop/public/images/c1.png)

![Graph 2](https://raw.githubusercontent.com/n3m/covid-dashboard-server/develop/public/images/c2.png)

![Graph 3](https://raw.githubusercontent.com/n3m/covid-dashboard-server/develop/public/images/c3.png)

![Graph 4](https://raw.githubusercontent.com/n3m/covid-dashboard-server/develop/public/images/c4.png)

![Graph 5](https://raw.githubusercontent.com/n3m/covid-dashboard-server/develop/public/images/c5.png)
