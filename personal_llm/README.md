#### Personal llm
A portable containerized llm hosted locally using [docker](https://www.docker.com/), [ollama](https://ollama.com/) and [openUI](https://docs.openwebui.com/). The llm can be ran on CPU and GPU enabled systems.

#### Requirements
- [Docker](https://www.docker.com/products/docker-desktop/)

#### Start steps
1. Start the docker daemon 
2. Navigate to project root directory and run ```docker-compose up -d```
3. Visit [localhost:8282](http://localhost:8282/) on a browser
4. Sign up for an account
5. Search and download [ollama models](https://ollama.com/library?sort=popular)
6. Select model and  prompt away

#### Stop steps
1. Navigate to project root directory and run ```docker-compose down --volumes``` to turn off containers and delete volumes

#### Resources
- Docker daemon start [guide](https://docs.docker.com/config/daemon/start/)

#### Troubleshooting
###### Windows operating system
- [Install Windows Subsystem for Linux(wsl)](https://learn.microsoft.com/en-us/windows/wsl/install)
- [Enable virtualization on windows](https://learn.microsoft.com/en-us/windows/wsl/troubleshooting#error-0x80370102-the-virtual-machine-could-not-be-started-because-a-required-feature-is-not-installed)

###### Mac operating system
- [Change docker desktop settings on mac](https://docs.docker.com/desktop/settings/mac/#namespaces)
- [Docker file sharing](https://docs.docker.com/desktop/settings/mac/?uuid=51156F3F-7CDF-494C-B5D6-B96B2060A073#file-sharing)