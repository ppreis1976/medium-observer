# Exemplo do Padrão Observer em Go para Sistema de Pedidos

Este projeto demonstra a implementação do padrão comportamental Observer em Go, utilizando um cenário de sistema de pedidos. O objetivo é ilustrar como o padrão Observer permite que diferentes partes de um sistema (neste caso, notificações por e-mail, WhatsApp e SMS) reajam a mudanças no estado de um objeto (um pedido) de forma desacoplada.

## Descrição

O padrão Observer define uma dependência um-para-muitos entre objetos, de forma que, quando um objeto (o sujeito) muda seu estado, todos os seus dependentes (os observadores) são notificados e atualizados automaticamente.

Neste exemplo, o `Pedido` atua como o sujeito, e os serviços de notificação por `Email`, `WhatsApp` e `SMS` atuam como observadores. Quando o status do pedido é atualizado, todos os observadores registrados são notificados para realizar a ação correspondente.

## Pré-requisitos

* **Go instalado:** Certifique-se de ter o Go instalado em sua máquina. Você pode baixar e instalar a partir do site oficial: [https://go.dev/dl/](https://go.dev/dl/)

## Como Executar

1.  **Clone o repositório (opcional):** Se este código estiver em um repositório Git, clone-o para sua máquina:
    ```bash
    git clone [URL_DO_REPOSITORIO]
    cd [NOME_DO_DIRETORIO]
    ```

2.  **Navegue até o diretório do projeto:**
    ```bash
    cd observer_pedidos
    ```
    (Assumindo que o arquivo principal `main.go` está dentro de um diretório chamado `observer_pedidos`).

3.  **Execute o programa:**
    ```bash
    go run main.go
    ```

## Saída Esperada

Ao executar o programa, você deverá ver uma saída semelhante a esta:

```
Pedido 123 teve seu status atualizado para: Processando
Enviando e-mail para cliente@email.com sobre o status do pedido 123: Processando
Enviando WhatsApp para 551199999999 sobre o status do pedido 123: Processando

Removendo notificação por WhatsApp...
Pedido 123 teve seu status atualizado para: Enviado
Enviando e-mail para cliente@email.com sobre o status do pedido 123: Enviado

Pedido 123 teve seu status atualizado para: Entregue
Enviando e-mail para cliente@email.com sobre o status do pedido 123: Entregue
Enviando SMS para 551199999999 sobre o status do pedido 123: Entregue
```

Esta saída demonstra o fluxo de notificações à medida que o status do pedido é atualizado e observadores são registrados e removidos.

## Explicação do Código

O projeto é composto pelos seguintes elementos principais:

* **`Observer` Interface:** Define o método `Update(order *Order)` que os observadores devem implementar para receber notificações sobre mudanças no pedido.
* **`ConcreteObserver`s:** Implementações concretas da interface `Observer`:
    * `EmailNotificationObserver`: Simula o envio de um e-mail.
    * `WhatsappNotificationObserver`: Simula o envio de uma mensagem via WhatsApp.
    * `SMSNotificationObserver`: Simula o envio de um SMS.
* **`Order` Struct:** Representa um pedido com informações como ID, status, e-mail e telefone do cliente. Possui um método `UpdateStatus` para alterar seu status.
* **`OrderSubject` Interface:** Define métodos para registrar (`RegisterObserver`), remover (`RemoveObserver`) e notificar (`NotifyObservers`) os observadores.
* **`ConcreteOrderSubject` Struct:** Implementação concreta da interface `OrderSubject`. Mantém uma lista de observadores e o estado do pedido. O método `UpdateOrderStatus` atualiza o status do pedido e notifica todos os observadores registrados.
* **`main.go`:** Contém a função principal que cria um pedido, instancia os observadores, os registra no sistema de pedidos e simula a atualização do status do pedido, demonstrando o padrão Observer em ação.
