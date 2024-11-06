# Mecânica de jogo

## Armas

As armas serão espalhadas pelo jogo de acordo com o nível em que o personagem se encontrar.

### Espada 🗡️

Amenta pontos de ataque e defesa (em menor quantidade). Disponível para:

- Humano
- Elfo
- Anão
- Mago

### Escudo 🛡️

Aumenta pontos de defesa. Disponível para:

- Humano
- Anão
- Elfo

### Arco e flecha 🏹

Aumenta pontos de ataque. Disponível para:

- Humano
- Elfo

### Machado 🪓

Aumenta pontos de ataque. Disponível para:

- Anão

### Cajado ⚕️

Aumenta pontos de ataque e defesa. Disponível para:

- Mago

### Varinha 🪄

Aumenta pontos de ataque e defesa. Disponível para:

- Humano
- Elfo
- Anão
- Mago

### Adagas 🗡️🗡️

Aumenta pontos de ataque. Disponível para:

- Elfo

## Armas pelos territórios

### Em Lund

Na terra de Lund o personagem poderá encontrar os seguintes itens:

- Espada serrilhada => +1 ponto na habilidade de ataque

- Machado de aço => +1 ponto na habilidade de ataque

- Escudo de aço => +2 ponto na habilidade de defesa

Ao enfrentar um oponente a chance de o personagem encontrar um dos itens acima é de 16.6%

### Em Horkjavic

Na terra de Horkjavic o personagem poderá encontrar os seguintes itens:

- Varinha dos quatro ventos => +2 pontos na habilidade de ataque e defesa

- Machado de Mithril => +3 pontos na habilidade de ataque

- Adagas de Mitrhil => +3 pontos na habilidade de ataque

- Escudo de Mithril => +4 pontos na habilidade de defesa

Ao enfrentar um oponente a chance de o personagem encontrar um dos itens acima é de 16.6%

### Em Ut'Manumbar

Na terra de Ut'Manumbar o personagem poderá encontrar os seguintes itens:

- Cajado de Ut'Gorithar => +5 pontos na habilidade de ataque e defesa

- Varinha das Sete Leis => +7 pontos na habilidade de ataque e defesa

- Arco de Mithril => +5 pontos na habilidade de ataque

- Machado de Foriahir => +5 pontos na habilidade de ataque

Ao enfrentar um oponente a chance de o personagem encontrar um dos itens acima é de 8.33%

## Habilidades das raças

As habilidades são o fator que aumenta a capacidade de ataque ou defesa do jogador, da seguinte forma:

- Mago = Sabedoria => Favorece a defesa: 3x pontuação
- Humano = Astúcia => Favorece o ataque: 2x pontuação
- Elfo = Quatro elementos => Favorece o ataque: 3x pontuação
- Anão = Forja => Favorece o ataque: 2x pontuação

## Níveis de dificuldade

Os níveis de dificuldade afetam na quantidade de inimigos enfrentados a cada terra do jogo.

### Em Lund:

- Nível Fácil: 3 orcs
- Nível Médio: 5 orcs
- Nível Difícil: 7 orcs

### Em Horkjavik

- Nível Fácil: 3 orcs + 1 feiticeiro
- Nível Médio: 5 orcs + 2 feiticeiros
- Nível Difícil: 7 + 3 feiticeiros

### Em Ut'Manumbar

- Nível Fácil: 3 orcs + 2 feiticeiros + 1 Lorde das Sombras + Colun
- Nível Médio: 5 orcs + 3 feiticeiros + 2 Lordes das Sombras + Colun
- Nível Difícil: 7 orcs + 4 feiticeiros + 3 Lordes das Sombras + Colun

## Poções

### Cura

As poções de cura serão espalhadas pelo jogo conforme o personagem enfrentar oponentes. Ao final de cada batalha, o personagem terá uma chance de 50% de encontrar uma poção de cura. A poção de cura age no personagem de acordo com a sua vitalidade. Curando-o com 50% de seu valor.

A fórmula a seguir define o valor final de uma ação de cura:

`vc = 50/100 * vv`

> Onde:
> vc = valor cura
> vv = valor da vitalidade do personagem

### Força

As poções de força são espalhadas pelo jogo conforme o personagem progredir de nível. Ao final de cada nível existe uma chance de 33.3% de encontrar uma poção de força.

A fórmula a seguir define o valor final de uma ação de aumento de força:

`vf = vf + 5`

> Onde:
> vf = valor força

## Fatores de ataque e defesa

A fórmula a seguir define o valor final de um ataque ou defesa:

`vf = f + vd * ph`

> Onde:
> vf = valor final de ataque/defesa
> f = força
> vd = valor dos dados
> ph = pontos de habilidade
