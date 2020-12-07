# Greenblatt Crawler

Este repositório contém uma aplicação escrita em Golang que busca ativos de empresas listadas na bolsa (B3) de acordo com determinados filtros. Com os dados computados, é aplicado uma fórmula personalizada baseada na de [Joel Greenblatt](https://www.amazon.com/Formula-Magica-Greenblatt-Mercado-Portugues/dp/8557173601?tag=techblast0f-20).

- Você consegue alterar os filtros listados no arquivo `filtros.yml`
- Os dados de saída serão gerados no arquivo `result.csv`

## Filtros

A aplicação permite os seguintes filtros:
- Dividendo
- P/L
- Dívida Líquida / EBIT
- Dívida Líquida / Patrimonio Liquido
- ROE
- ROIC
- ROA
- Liquidez Média Diária

No arquivo de configuração, você deve seguir o seguinte padrão:

```yaml
filtros:
  dividendo: (em porcentagem)
    - valor mínimo (ex: 0.1)
    - valor máximo (ex: 100)
  liqmeddiaria:
    - 1000000
    - 10000000000
```

## Ordenação

A ordenação aplicada segue os seguintes criterios:

- P/L ordem crescente
- ROIC ordem decrescente
- Dividendo ordem decrescente
- Posição final -> soma das posições das ordenações anteriores

## Buildando aplicação

### Sem Docker

Se você possui Golang instalado com a versão 1.11.x ou maior, execute o comando

```bash
make build
```

### Com Docker

Se você prefere buildar a aplicação utilizando um container docker, execute o comando

```bash
make dkbuild
```

### Rodando aplicação

Após a etapa de build, será gerado um arquivo na pasta `build/greenblatt_crawler(.exe)`. Execute-o apenas clicando (Windows) ou da seguinte forma em ambientes Unix

```bash
./build/greenblatt_crawler
```

Saída esperada

```
Carregando arquivo contendo os filtros...
Buscando recursos na API Status Invest...
Ordenando rank de companhias...
Salvando dados em arquivo...
Prontinho! Seus dados estão no arquivo result.csv
```

Seja feliz!
