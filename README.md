- 4 web apps that are consumer applications of the eslint config, tsconfig and ui library 
- Was inteneded as a great showcase for RabbitMQ, but turned out to be a great example of monorepo usage as well
- The communication between the 4 frontends and the server is an representing a real world challange of multi regional ticket vendors 
- Go server with Mux 
- The 4 frontends are AMQP producer, where the Go server is an AMQP consumer
- When a user clicks one of the buttons, the frontend application generates a message. This message contains the details about the number of tickets being bought (+10 or +100).
- The Go server acts as a consumer in AMQP, an entity that connects to the message broker to receive messages and process them
- RabbitMQ as a Broker inbetween
- After processing the message, the server sends a confirmation back to the initial site (the frontend application) through RabbitMQ.

- I wanted the repo to be able to run withouth complete dockerization, with mostly turbo running everything except the rabbitmq image that is needed for the server

- The go server serves as a gateway service and a processing service at the same time as its a poc project (in real life project this would be decoupled)

4 Web Apps for Ticket Booking:
- Each app represents a consumer application for different countries (France, Germany, Great Britain, Spain).
- The applications use shared configuration and UI components, demonstrating efficient code reuse in a monorepo setup.

Monorepo with Turborepo:
- The project showcases an effective monorepo setup using Turborepo for managing multiple interconnected projects.
- The monorepo contains common configurations (ESLint, TypeScript), shared UI components, and individual applications for each country.

RabbitMQ for Messaging:
- RabbitMQ serves as the message broker between the frontends and the server, representing a real-world multi-regional ticket vendor system.
-The communication flow demonstrates an example of AMQP usage, with the frontends acting as producers and the server as a consumer.

Go Server with Mux:
- The server is implemented in Go using the Gorilla Mux router for handling HTTP requests.
- It processes messages from the RabbitMQ broker and sends confirmations back to the relevant frontend.

Live Reload Development: Configured the Go server for live reloading using Air, enhancing developer experience by enabling immediate feedback during development.

Consumer Setup:
- RabbitMQ queues are declared during server startup, ensuring necessary infrastructure is set up before processing begins.
- Utilizes concurrent processing through goroutines for efficient consumer initialization.
- The server processes messages based on the country code provided, and dispatches them to different queues accordingly.

Handling Ticket Purchases:
- When a user clicks one of the buttons on the frontend, a message is generated detailing the number of tickets being bought.
- The server processes these messages and sends confirmations back to the appropriate frontend.

Data Flow:
- The frontends send API requests to the Go server.
- The server validates and processes these requests, forwarding the messages to RabbitMQ.
- RabbitMQ distributes the messages to the appropriate queues for each country.

Fairness in Message Handling:
- The RabbitMQ broker ensures that ticket purchase requests are fairly distributed among the different countries, avoiding one country monopolizing the tickets.


// Todo:

Confirmation Workflow:
After processing the message, the server sends a confirmation message back to the originating frontend application.


The confirmation is sent through RabbitMQ, maintaining the asynchronous, decoupled communication flow.


That's what I'm working on now:
Update order status in your data store and send updates
			// updateOrderStatus(p)
		// Broadcast the updated status to the client


Error Handling:
- The server includes robust error handling, providing clear feedback and maintaining consistent logging for monitoring purposes.
Logging and Monitoring:
- The Go server logs key actions and errors, aiding in debugging and monitoring the application’s health.



## ChatGPT anser about RabbitMq Portfolio Project:

Yes, RabbitMQ is an excellent choice for your project to handle ticket orders coming from multiple websites. Using RabbitMQ as a message broker in this scenario offers several key benefits:

    ### Decoupling: RabbitMQ will effectively decouple your websites from the backend processing system. This means that the websites can quickly send order messages to the queue without having to wait for the order processing and confirmation, enhancing the responsiveness of your websites.

    ### Load Management: RabbitMQ can help in managing load by queueing incoming ticket orders and processing them one at a time or in small batches. This prevents overloading your server when there are spikes in traffic, such as during popular events or promotions.

    ### Reliability: RabbitMQ ensures that once an order is placed into the queue, it won’t be lost. Even if the backend system fails or needs to be restarted, the messages will be retained and can be processed once the system is back online.

    ### Scalability: As your traffic grows or during peak times, RabbitMQ facilitates scaling your processing workers horizontally. You can add more consumers to the queues to handle increased loads without any changes to the websites or the core backend logic.

    ### Asynchronous Communication: It allows the backend processing to be handled asynchronously. This means that your websites don’t need to hold connections open while the backend processes the orders, which can improve the overall efficiency and user experience.

    ### Order Integrity and Sequencing: If needed, RabbitMQ can help maintain the order of messages. This is crucial for ticketing systems where the sequence of ticket sales might affect availability and booking confirmations.


# Turborepo starter

This is an official starter Turborepo.

## Using this example

Run the following command:

```sh
npx create-turbo@latest
```

## What's inside?

This Turborepo includes the following packages/apps:

### Apps and Packages

- `docs`: a [Next.js](https://nextjs.org/) app
- `web`: another [Next.js](https://nextjs.org/) app
- `@repo/ui`: a stub React component library shared by both `web` and `docs` applications
- `@repo/eslint-config`: `eslint` configurations (includes `eslint-config-next` and `eslint-config-prettier`)
- `@repo/typescript-config`: `tsconfig.json`s used throughout the monorepo

Each package/app is 100% [TypeScript](https://www.typescriptlang.org/).

### Utilities

This Turborepo has some additional tools already setup for you:

- [TypeScript](https://www.typescriptlang.org/) for static type checking
- [ESLint](https://eslint.org/) for code linting
- [Prettier](https://prettier.io) for code formatting

### Build

To build all apps and packages, run the following command:

```
cd my-turborepo
pnpm build
```

### Develop

To develop all apps and packages, run the following command:

```
cd my-turborepo
pnpm dev
```

### Remote Caching

Turborepo can use a technique known as [Remote Caching](https://turbo.build/repo/docs/core-concepts/remote-caching) to share cache artifacts across machines, enabling you to share build caches with your team and CI/CD pipelines.

By default, Turborepo will cache locally. To enable Remote Caching you will need an account with Vercel. If you don't have an account you can [create one](https://vercel.com/signup), then enter the following commands:

```
cd my-turborepo
npx turbo login
```

This will authenticate the Turborepo CLI with your [Vercel account](https://vercel.com/docs/concepts/personal-accounts/overview).

Next, you can link your Turborepo to your Remote Cache by running the following command from the root of your Turborepo:

```
npx turbo link
```

## Useful Links

Learn more about the power of Turborepo:

- [Tasks](https://turbo.build/repo/docs/core-concepts/monorepos/running-tasks)
- [Caching](https://turbo.build/repo/docs/core-concepts/caching)
- [Remote Caching](https://turbo.build/repo/docs/core-concepts/remote-caching)
- [Filtering](https://turbo.build/repo/docs/core-concepts/monorepos/filtering)
- [Configuration Options](https://turbo.build/repo/docs/reference/configuration)
- [CLI Usage](https://turbo.build/repo/docs/reference/command-line-reference)
