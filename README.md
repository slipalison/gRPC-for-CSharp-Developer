# Go Lang
O GO tem uma particularidade referente as bibliotecas, todo seu repositorio deve estar na pasta **$GOPATH\src**...
Configurar a variavel de ambiente GOPATH (o padrão no windows %USERPROFILE%\go) para sua pasta de projetos padão;

> EX: C:\projetos\go

Quando se trabalha com GIT, clonar o repositório dentro da pasta:
> $GOPATH\src\\[MeuDominio]\

Reinicie o PC se necessário.

### Adicionando DEP
Para instalar o Dep na sua maquina basta ter o Go 1.11 ou superior em sua maquina e executar o seguinte comando:
```
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```
E então esse comando para instalar as dependencias:
```
    dep init                               set up a new project
    dep ensure                             install the project's dependencies
    dep ensure -update                     update the locked versions of all dependencies
    dep ensure -add github.com/pkg/errors  add a dependency to the project
```
### Testes
Para rodar todos os testes e coletar a cobertura em html:
```
 go test ./... -coverprofile=/app/coverage.out && go tool cover -html=/app/coverage.out -o=/app/coverage.html
 ```

# PROTOBUF
### Quando utilizar?
É importante observar que, embora as mensagens JSON e  _Protobuf_ possam ser usadas de forma intercambiável, essas tecnologias foram projetadas com objetivos diferentes. JSON, que significa  _JavaScript Object Notation_, é simplesmente um formato de mensagem que surgiu em subconjunto da linguagem de programação  _Javascript._ As mensagens JSON são trocadas em formato de texto e, hoje em dia, são complemente independentes e suportadas por praticamente todas as linguagens de programação.

_Protobuf_, por outro lado, é mais que um formato de mensagem, é também um conjunto de regras e ferramentas para definir e trocar essas mensagens. Além disso, o  _Probuf_ possui mais tipos de dados que o JSON, como enumerados e métodos, e também é muito usado em  [RPCs( Remote Procedure Calls).](https://github.com/grpc)

Abaixo cito cinco boas razões para utilizar  _Protobuf:_

1.  **Schemas.** Existe toda uma preocupação em normalizar os dados e gravar eles corretamente nas nossas bases de dados. Contudo, por quê não existe tal cuidado quando transportamos nossos dados usando serviços de mensagem? O  _Protobuf_ resolve esta limitação permitindo que sejam definidos esquemas de dados que definem e garantem a integridade dos dados transportados entre diferentes serviços.

2.  **Compatibilidade de versões anteriores.** A maneira que os campos são definidos no  _schema_( de maneira enumerada),  evita a necessidade de verificação de compatibilidade entre versões, que é uma das principais razões do  _Protobuf_ ter sido projetado.

3.  **Menos código duplicado.**  Geralmente os  _endpoints_  HTTP que se comunicam via JSON dependem de bibliotecas de terceiros ou códigos implementados especificamente para codificar e decodificar o JSON. Além do mais, muitas vezes as classes referentes ao JSON acabam expondo regras de negócio e dificultando ainda mais o trabalho manual de parser. O Protobuf resolve este problema ao gerar classes  _stub_ a partir dos  _schemas_ definidos. Deste modo,conforme os dados forem evoluindo e mudando, será somente necessário modificar o  _schema_ e regerar as classes  _stub_.

4.  **Validações e extensibilidade**. As palavras chaves  `required` ,  `optional`  e  `repetead`  são extremamente poderosos. Elas permitem que seja codificado, a nível de  _schema_, o formato dos dados estruturados e os detalhes de implementação sobre como as classes devem funcionar. Por exemplo, a biblioteca de  _Protobuf_ do  _Python_ irá mostrar uma exceção na tentativa de instanciar um objeto que possua um campo  `required`  vazio. É possível também alterar um campo de  `required`  para  `optional` — ou vice-versa — simplesmente trocando a enumeração do campo. Conforme o projeto vai crescendo e surge a necessidade de modificar a estrutura de dados, este tipo de flexibilidade realmente se mostra importante.

5.  **Interoperabilidade entre linguagens.**  Como os  _Protocol Buffers_  são implementados em várias linguagens, eles tornam a interoperabilidade entre aplicativos poliglotas em sua arquitetura muito mais simples. Independente da linguagem utilizada no serviço que for adicionado na arquitetura, utilizando o  _Protobuf_ somente é preciso ter em mãos o arquivo `.proto`  e gerar a classe  _stub_  da linguagem do novo serviço.

### E quando JSON é melhor?

Ainda há algumas vezes que o JSON é melhor que o  _Protobuf_, como por exemplo nas situações em que:

-   É necessário que os dados sejam legíveis para humanos.
-   Os dados do serviço são consumidos diretamente por um  _web browser_.
-   Sua aplicação  _server side_ é escrita em  _JavaScript_.
-   Você não está preparado para vincular o modelo de dados a um esquema, por exemplo, talvez seus dados são dinâmicos.
-   A sua aplicação não consome tanta banda assim.
[Protobuf — Uma alternativa ao JSON e XML](https://medium.com/trainingcenter/protobuf-uma-alternativa-ao-json-e-xml-a35c66edab4d)



# gRPC
Grandes empresas como Square, Netflix, Digital Ocean e SoundCloud estão adotando o gRPC para fazer seus microserviços se comunicarem melhor.

Mas o que é esse tal de gRPC?

Primeiramente é importante definirmos o que é RPC.

Uma chamada de procedimento remoto (RCP) se trata de uma tecnologia para comunicação de processos, permitindo que um programa de computador seja capaz de chamar um procedimento que está em outro computador, independentemente da linguagem e plataforma. O gRPC é um framework opensource RPC feito pelo Google.

A principal diferença entre ele e o REST é que o segundo é focado em recursos, enquanto o RPC é focado em operações. O exemplo abaixo pode ilustrar isso melhor:

|Ação            |RPC(operação)                  |REST(recurso)                |
|----------------|-------------------------------|-----------------------------|
|Signup			 |`POST /signup`                 |POST /pessoas                |
|Resign          |`POST /resign`                 |DELETE /pessoas/1234         |
|Ler uma pessoa  |`GET /getPessoa?personid=1234` |GET /pessoas/1234            |
|Atualizar pessoa|`post /updateperson`           |PUT /pessoas/1234            |

O GRPC se aplica bem e é altamente recomendado em cenários em que são esperados:

-   Baixa latência;
-   Alta escalabilidade;
-   Sistemas distribuídos;
-   Aplicativos mobile que se comunicam com servidores na nuvem;

Isso porque ele é construído sobre o HTTP2 ([**clique aqui para um teste de velocidade comparado ao HTTP1**](http://www.http2demo.io/)) e além de fornecer ferramentas eficientes de autenticação, propicia balanceamento de carga, monitoramento, logging etc.

Aqui podemos observar qual é realmente o lugar que ele ocupa em uma arquitetura de microserviços.

![enter image description here](https://cdn-images-1.medium.com/max/800/0*ht0MhpWv8CNP__Sz.)

Essa comunicação entre o serviço C e o restante do sistema está sendo feita por gRPC e o formato de dados é o  [**_Protobuf_**](https://developers.google.com/protocol-buffers/). Pense em um XML mais rápido, simples e menor, isso é o Protobuf! Sua rapidez é devido ao seu encoding binário.


Em redes de comunicação, como Ethernet ou packet radio, **throughput**,**throughput** de rede ou simplesmente taxa de transferência é a quantidade de dados transferidos de um lugar a outro, ou a quantidade de dados processados em um determinado espaço de tempo.

Um comparativo entre o desempenho de JSON/REST e Protobuf/gRPC:

![enter image description here](https://cdn-images-1.medium.com/max/800/0*D5M7X-pYXGHNGhCU.)

Claro que nem tudo são flores e em algumas situações AINDA não é aconselhável trocar o JSON por Protobuf. São eles elas:

-   Quando precisamos ter dados legíveis para humanos;
-   Quando os dados são consumidos diretamente pelo navegador;
-   Seu lado servidor é escrito em JavaScript

O gRPC fornece uma  [documentação](http://www.grpc.io/docs/)  bastante legal e o guia básico pode ser feito em várias linguagens, dentre elas Java, C#, PHP… e por aí vai! [exemplos](https://github.com/grpc/grpc).
[Melhorando o desempenho de microservices com gRPC](https://medium.com/quick-mobile/melhorando-o-desempenho-de-microservices-com-grpc-31bd67d210e7)

