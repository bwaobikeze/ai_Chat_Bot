#### Personal llm
A portable CPU based llm using docker and ollama. 

#### Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)

#### Start
1. Start the docker daemon
    - Docker daemon start [guide](https://docs.docker.com/config/daemon/start/) 
2. Navigate to project root directory and run ```docker-compose up -d```
3. Visit [localhost:8282](http://localhost:8282/) on a browser

#### Stop
1. Navigate to project root directory and run ```docker-compose down --volumes```

#### Resources
- [ollama models](https://ollama.com/library?sort=popular)

#### Windows trouble shoot
- [Install Windows Subsystem for Linux(wsl)](https://learn.microsoft.com/en-us/windows/wsl/install)
- [Enable virtualization on windows](https://learn.microsoft.com/en-us/windows/wsl/troubleshooting#error-0x80370102-the-virtual-machine-could-not-be-started-because-a-required-feature-is-not-installed)

#### Mac trouble shooting
- [Change docker desktop settings on mac](https://docs.docker.com/desktop/settings/mac/#namespaces)
- [Docker file sharing](https://docs.docker.com/desktop/settings/mac/?uuid=51156F3F-7CDF-494C-B5D6-B96B2060A073#file-sharing)