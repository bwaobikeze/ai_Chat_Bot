#### Personal llm
A portable CPU based llm using docker and ollama.

#### Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)
- [ollama](https://ollama.com/download/)

#### Start
1. Start the docker daemon
    - Docker daemon start [guide](https://docs.docker.com/config/daemon/start/) 
2. Navigate to project root directory and run ```docker-compose up -d```
3. Visit [localhost:8282](http://localhost:8282/) on a browser

#### Resources
- [ollama models](https://ollama.com/library?sort=popular)

#### Windows trouble shoot
- [Install Windows Subsystem for Linux(wsl)](https://learn.microsoft.com/en-us/windows/wsl/install)
- [Enable virtualization on windows](https://learn.microsoft.com/en-us/windows/wsl/troubleshooting#error-0x80370102-the-virtual-machine-could-not-be-started-because-a-required-feature-is-not-installed)