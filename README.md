# Goroutines na prática

Esse simples projeto foi feito para entender como goroutines funcionam de uma forma mais lúdica e fácil.

## Conhecimentos Aplicados

- Goroutines
- Channels
- WaitGroups

## Goroutines

Imagine que você trabalha na logística de uma empresa que envia produtos para clientes. Sua empresa está bastante lenta, pois só tem um funcionário para embalar todos os produtos. Seu funcionário leva 10 minutos para embalar um único produto. Então, se tiver 6 produtos, ele levará uma hora.

Você então decide aplicar a prática de goroutines e contrata um funcionário para cada produto a ser embalado e por um valor muito baixo. Você percebe que agora seus produtos estão sendo embalados muito mais rápido que antes e seus clientes estão muito mais felizes.

As goroutines são funções que são executadas simultaneamente em tempo de execução, com baixo custo, acelerando todo o seu código.

## Channels

Agora que seus funcionários (goroutines) estão embalando os produtos ao mesmo tempo, eles precisam ter uma esteira para enviar o produto para o mesmo caminhão de entrega. Essa esteira são os channels em Go.

Os canais servem para que as goroutines enviem dados para algum lugar do seu código. Todas as goroutines de uma função vão usar o mesmo canal para enviar os dados.

Existem mais detalhes sobre channels, como canais bufferizados e deadlocks, que valem a pena ser vistos.

## WaitGroups

Imagine que você trabalha na área de logística de uma empresa e precisa enviar um caminhão de produtos para o seu cliente. Para fazer esses produtos, você tem uma série de funcionários (goroutines), cada um fazendo um produto ao mesmo tempo.

Você percebe que o seu caminhão está sendo enviado a cada produto feito por um funcionário e isso está atrasando bastante o fluxo de trabalho. Daí você descobre algo chamado WaitGroups.

Os WaitGroups são como um supervisor que só libera o caminhão quando todos os produtos são finalizados. Ou seja, WaitGroups são uma maneira de você "travar" o andamento do seu código até que todas as goroutines desse WaitGroup sejam finalizadas.