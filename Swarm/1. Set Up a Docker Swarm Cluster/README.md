# Objective: Set Up a Docker Swarm Cluster

**Tasks**:
    1. Initialize a cluster of Docker Engines in swarm
    2. Add nodes to the Swarm
    3. Deploy application services to the Swarm
    4. Manage the Swarm once you have everything running

1. ## Initialize a cluster of Docker Engines in swarm

    To demostrate we'll need 3 Linux Hosts each having Docker installed and will communicate over a network. For this task, we'll be using Microsoft's Azure CLI to create the VMs (Virtual Machines).

    ======================= MS Azure VMs ======================== 

    [Visual Aid showing created VMs on Azure]
    ### A. Steps in Creating 3 VMs using Azure CLI

        1. Login to your azure account.
            $ az login
        
        2. Create a resource group.
    
            $ az group create --name swarm_res_group --location southafricanorth

        3. Create a Virtual Network.

            $ az network vnet create --name swarm_vnet --resource-group swarm_res_group --subnet-name swarm_subnet --subnet-range 10.0.0.0/24
        
        4. Create a public IP address for each VM.

            $ az network public-ip create --name vm1p_ip --resource-group swarm_res_group
            $ az network public-ip create --name vm2p_ip --resource-group swarm_res_group
            $ az network public-ip create --name vm3p_ip --resource-group swarm_res_group

        5. Create a network interface for each Virtual Machine.

            $ az network nic create --name vm1_nic --resource-group swarm_res_group --vnet-name swarm_vnet --subnet swarm_subnet --public-ip-address vm1p_ip
            $ az network nic create --name vm2_nic --resource-group swarm_res_group --vnet-name swarm_vnet --subnet swarm_subnet --public-ip-address vm2p_ip
            $ az network nic create --name vm3_nic --resource-group swarm_res_group --vnet-name swarm_vnet --subnet swarm_subnet --public-ip-address vm3p_ip
        
        6. Create the Virtual Machines

            $ az vm create --resource-group swarm_res_group --name node1 --image Ubuntu2204 --size Standard_D2s_v2 --nics vm1_nic --admin-username azureuser
            $ az vm create --resource-group swarm_res_group --name node2 --image Ubuntu2204 --size Standard_D2s_v2 --nics vm2_nic --admin-username azureuser
            $ az vm create --resource-group swarm_res_group --name node3 --image Ubuntu2204 --size Standard_D2s_v2 --nics vm3_nic --admin-username azureuser

        **Paste a screenshot of the created VMs from Azure Portal**

2. ## Add nodes (VMs) to the Swarm

    [Commands and Visual Aid]

3. ## Deploy application services to the Swarm

    [Commands and Visual Aid]


4. ## Manage the Swarm once you have everything running

    [Commands and visual aid]