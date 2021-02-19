# ClubTree
Make your own "invitation tree graph" of your favorite invitational SNS like `Clubhouse`.

## Usage

```
$ go main.go print -j samples/tree.json
/ 
└── max/Max Planck
│   ├── niels/Niels Bohr
│       └── werner/Werner Heisenberg
│       └── erwin/Erwin Schrodinger
└── paul/Paul Dirac
$ go run main.go add -j samples/tree.json
Enter the clubhouse ID
einstein
Enter the alias

Enter the first name
Albert
Enter the last name
Einstein
Enter the Twitter ID

Enter the Instagram ID

Enter the nominator ID
alfred
$ go run main.go print -j samples/tree.json
/
└── max/Max Planck
│   ├── niels/Niels Bohr
│       └── werner/Werner Heisenberg
│       └── erwin/Erwin Schrodinger
└── paul/Paul Dirac
└── einstein/Albert Einstein
```
