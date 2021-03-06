# Droid Battles | Ludum Dare 50

Droids were created initially to help us with daily tasks, like cleaning up the
house or driving us to work. Today, they are also used to build up Droid Teams
and compete against each other in the Droid Championship.

You just earned the Droid Builder License and can now compete as well and go for
a chance to change your family luck once and for all. As a gift, your parents,
you start with Blue, the family outdated but functional cleaning droid.

You don't have enough Credits to buy a proffesional Battle Droid, so you have to
try it out with the unoffial and dangeoours battles first and build up your
skills and your Droid from there.

## Task List

- [ ] Game Mechanics
    - [x] Start Game with title screen
        - [x] Exit with ESC key
    - [x] Show player inventory
        - [x] Show inventory screen with ENTER key
        - [x] Go back from inventory to Title when ESC key is pressed
        - [x] Show all droids the player has unlocked
        - [x] Show all chips the player has unlocked
    - [ ] Battle Mecanics
        - [ ] PvE Mode (Unit upgrades, story mode)
            - [ ] First phase - 1x1 unofficial battles
                - [x] 1 character for each party
                - [x] 3 skills to use each round
                - [ ] 10 rounds limit per battle
        - [ ] PvP Mode (Arena Battles)
            - [ ] Second phase - 3x3 championship battles
                - [ ] 3 characters for each party
                - [ ] Passive skills from all party members
                - [ ] 20 rounds limit per battle
    - [ ] Character evolution
        - [x] Blue starting droid give to player
        - [ ] Droids can upgrade by adding up to 6 Chips (runes)
        - [ ] Chips improve stats and the overall Droid power
        - [ ] Each unit has 3 optional skills and 1 passive
        - [ ] Each skill has a cooldown that requires eletricity to trigger
- [ ] Artwork
    - [x] Game Title Screen
    - [x] Each droid is represented as a unit in the player inventory
    - [x] Droids will have improving look as they are upgraded
    - [ ] Droid rarity is visible as the droid portrait decoration
    - [ ] Pixel Art or 3D Art? (Each skill move will be the charm of the game, nice animations effects)
- [ ] Sound
    - [ ] Droid sounds will be motors, lasers and metal sfx
    - [x] BGM will be techno during battles, easy going on game screns
        - [x] Intro BGM
        - [x] Battle BGM
- [ ] Infra
    - [x] Hosting (Cloud Run)
    - [x] Automate deployment using Github Actions
    - [ ] Cross platform build

## Testing out

Install dependencies

    sudo apt install -yq libc6-dev libglu1-mesa-dev libgl1-mesa-dev \
        libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev \
        libxxf86vm-dev libasound2-dev librsvg2-bin pkg-config 

    go run github.com/ronoaldo/ld-50@latest

## Updating assets

After making changes to the assets/*.svg files, it is needed to update the
corresponding PNG ones. This can be done using `go generate`, but it is required
that you have `librsvg2-bin` installed.

    go generate ./...