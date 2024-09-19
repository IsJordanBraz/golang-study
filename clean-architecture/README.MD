- Hexagonal Architecture
- Onion Architecture

# Clean Architecture

* Proteção do domínio da aplicação.
* Baixo acoplamento entre as camadas.
* Orientada a casos de uso.

- Objetivo principal da arquitetura é dar suporte ao ciclo de vida do sistema.
- Uma boa arquitetura torna o sistema fácil de entender, Fácil de desenvolver, fácil de manter e fácil de implantar.
- O objetivo final é minimizar o custo de vida útil do sistema e maximizar a produtividade do programador.

### Regras vs detalhes

* (frameworks, banco de dados, apis, não devem impactar as regras).
* Regras de negócio trazem o real valor para o software.
* Detalhes ajudam a suportar as regras
* Detalhes não devem impactar nas regras de negócio 

### Use cases - Single Responsibility Principle

* Intenção, cada ação é uma intenção, cada intenção é um caso de uso.
* Clareza de cada comportamento do software
* Detalhes não devem impactar nas regras de negócio.
* "Reaproveitar" use cases por serem parecidos, são casos de usos diferentes.

### Input vs Output

- Cria um pedido vs Dados do pedido

### Presenters

- Objetos de transformação
- adequa o DTo de output no formato correto para entregar o resultado.
- Ex: XML, JSON, Protobuf, graphql, cli

### Entities

- Entities da clean != Entities do DDD
- Entity camada de regras de negócio
- Normalmente utiliza táticas do DDD
- Entites = agregados + domain Services

