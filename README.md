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
        - [ ] Start game with ENTER key
    - [ ] Show player inventory
    - [ ] Battle Mecanics
    - [ ] PvE Mode (Unit upgrades, story mode)
        - [ ] First phase - 1x1 unofficial battles
            - [ ] 1 character for each party
            - [ ] 3 skills to use each round
            - [ ] 10 rounds limit per battle
        - [ ] Second phase - 3x3 championship battles
            - [ ] 3 characters for each party
            - [ ] Passive skills from all party members
            - [ ] 20 rounds limit per battle
        - [ ] Character evolution
            - [ ] Blue starting droid give to player
            - [ ] Droids can upgrade by adding up to 6 Chips (runes)
            - [ ] Chips improve stats and the overall Droid power
            - [ ] Each unit has 3 optional skills and 1 passive
            - [ ] Each skill has a cooldown that requires eletricity to trigger
- [ ] Artwork
    - [ ] Game Title Screen
    - [ ] Each droid is represented as a unit in the player inventory
    - [ ] Droids will have improving look as they are upgraded
    - [ ] Droid rarity is visible as the droid portrait decoration
    - [ ] Pixel Art or 3D Art? (Each skill move will be the charm of the game, nice animations effects)
- [ ] Sound
    - [ ] Droid sounds will be motors, lasers and metal sfx
    - [ ] BGM will be techno during battles, easy going on game screns

## Testing out

    go run github.com/ronoaldo/ld-50@latest

## Updating assets

After making changes to the assets/*.svg files, it is needed to update the
corresponding PNG ones. This can be done using `go generate`, but it is required
that you have `librsvg2-bin` installed.

    sudo apt-get install -yq librsvg2-bin
    go generate ./...