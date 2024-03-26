# smite-mono
This is a mono repo for the microservices that make up the smite.one backend. Each service runs in a docker container and communicates with gRPC or http requests. This was my first time building microservices, I had a lot of fun with it.

In January 2024 Smite 2 was announced and requests for HiRez/Smite's developer API were closed, effectively killing the project. However, I perservered and made a mock-api to try and have a mostly working minimum viable product if they ever open their applications back up.

This is the backend, you can find the frontend here: https://github.com/undo-k/smite-app

Instructions for running:
1. Clone both repos
2. `cd smite-mono`
3. `make up_build`
4. `cd .. && cd smite-app`
5. `yarn dev` or `yarn dev --host` if you want to view it on mobile

The site should then be running on http://localhost:3000/

