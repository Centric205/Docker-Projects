# Objective: Set Up a Docker Swarm Cluster

**Tasks**:
    1. Initialize a cluster of Docker Engines in swarm
    2. Add nodes to the Swarm
    3. Deploy application services to the Swarm
    4. Manage the Swarm once you have everything running

1. ## Initialize a cluster of Docker Engines in swarm

    To demostrate I'll need 3 Linux Hosts each having Docker installed and will communicate over a network. For this task, I'll be using Canonical's Multipass to create Ubuntu VMs (Virtual Machines).

    Using **Multipass**, I have created 3 VMs as shown below:

    ![Alt text](3-VMs.png)


2. ## Add nodes (VMs) to the Swarm

    In the **ManagerNode** I typed the following to initialize the Swarm:

        $ docker swarm init

    This was the result:

    ![Alt text](image.png)

    I used the highlighted command to join **WorkerNode** and **WorkerNode2** into the swarm;

    **NOTE**: If the highlighted command above is not available, then type the following:

        $ docker swarm join-token worker

    ![Alt text](image-1.png)

    We now have a 3-Node Swarm.

3. ## Deploy application services to the Swarm

    Now I'll be deploying a **service** to the Swarm bu running the command below in the **MananagerNode**:

        $ docker service create --replicas 1 --name helloworld alpine ping docker.com

        ![Alt text](image-2.png)


4. ## Manage the Swarm once you have everything running

    [Commands and visual aid]