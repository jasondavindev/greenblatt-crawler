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

```
dy
p_L
p_VP
p_Ativo
margemBruta
margemEbit
margemLiquida
p_Ebit
eV_Ebit
dividaLiquidaEbit
dividaliquidaPatrimonioLiquido
p_SR
p_CapitalGiro
p_AtivoCirculante
roe
roic
roa
liquidezCorrente
pl_Ativo
passivo_Ativo
giroAtivos
receitas_Cagr5
lucros_Cagr5
liquidezMediaDiaria
vpa
lpa
```

### Filtros extras

A API permite filtrar por alguns filtros extras que obecedem um formato diferente. Alguns filtros:

| Nome | Tipo | Valor padrão |
| ---- | ---- | ------------ |
| Sector | string | vazio |
| SubSector | string | vazio |
| Segment | string | vazio |
| my_range | string | "0;25" |

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
