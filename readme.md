# Greenblatt Crawler

Este repositório contém uma aplicação escrita em Golang que busca ativos de empresas listadas na bolsa (B3) de acordo com determinados filtros. Com os dados computados, é aplicado uma fórmula personalizada baseada na de [Joel Greenblatt](https://www.amazon.com/Formula-Magica-Greenblatt-Mercado-Portugues/dp/8557173601?tag=techblast0f-20).

- Você consegue alterar os filtros listados no arquivo `filtros.yml`
- Os dados de saída serão gerados no arquivo `result.csv`

## Filtros

No arquivo de configuração, você deve seguir o seguinte padrão:

```yaml
filtros:
  dy: # (em porcentagem)
    - valor mínimo # (ex: 0.1)
    - valor máximo # (ex: 100)
  roe:
    - 0.1 # (quando emitido o segundo valor, o valor padrão será zero (0))
```

Filtros disponíveis:

- dy
- p_L
- p_VP
- p_Ativo
- margemBruta
- margemEbit
- margemLiquida
- p_Ebit
- eV_Ebit
- dividaLiquidaEbit
- dividaliquidaPatrimonioLiquido
- p_SR
- p_CapitalGiro
- p_AtivoCirculante
- roe
- roic
- roa
- liquidezCorrente
- pl_Ativo
- passivo_Ativo
- giroAtivos
- receitas_Cagr5
- lucros_Cagr5
- liquidezMediaDiaria
- vpa
- lpa

### Filtros extras

A API permite filtrar por alguns filtros extras que obecedem um formato diferente. Alguns filtros:

| Nome | Tipo | Valor padrão |
| ---- | ---- | ------------ |
| Sector | string | vazio |
| SubSector | string | vazio |
| Segment | string | vazio |
| my_range | string | "0;25" |

## Agrupamento

É possível fazer uma busca agrupando por areas e o Top N de cada uma das areas. As areas são `segment`, `sector` e `subsector`. Para cada busca por area, há uma pré-ordenação, assim, garantindo a ordenação correta do resultado final.

### Valores disponíveis

- sector
- subsector
- segment

### Arquivo de configuração

```yaml
filtros:
  ...
  ...
agrupamento:
  top: 10 # Os 10 primeiros no rank
  class: "sector"
```

Se as configurações de agrupamento forem omitidas, a busca padrão será executada (não há agrupamento). Se o valor `top` for omitido, o valor `10` será usado como padrão.

## Campos de saída

Campos de saída são as colunas que serão adicionadas ao arquivo de saída. Por exemplo, nome da companhia, ticker, ROE, Preço etc.

### Valores disponíveis

- Companhia
- Ticker
- Preco
- PL
- Dividendo
- Roic
- DividaLiquidaEbit
- DividaliquidaPatrimonioLiquido
- EVEbit
- LiquidezMediaDiaria
- Lpa
- Roe
- LucrosCagr5
- ReceitasCagr5
- MargemBruta
- MargemEbit
- MargemLiquida
- PEbit
- PVP
- Roa
- Vpa

Há também um campo extra `Area` que só é disponibilizado caso você utilize a busca por agrupamento.

### Arquivo de configuração

```yaml
filtros:
  ...
  ...

agrupamento:
  top: 10
  class: "sector"

campos_saida:
  - Companhia
  - Ticker
  - Area # Somente nas busca com agrupamento habilitado
  - Preco
  - PL
```

**As colunas respeitarão a ordem do campos listados.**

## Ordenação

A ordenação aplicada segue os seguintes criterios:

| Nome | Regra |
| ---- | ----- |
| P/L  | Ordem crescente |
| ROIC | Ordem decrescente |
| Dividendo | Ordem decrescente |
| Posição final | Soma das posições das ordenações anteriores |

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
./bin/greenblatt_crawler
```
