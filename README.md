# tile-puzzle

## About the project
This project is for learning purpose only. It's an implementation of the wellknown tile puzzle game in golang.
The goal is to solve the game by moving the space tile around with the WASD keys. The game is over if you're
able to create the following state:

```bash
1 2 3
4 5 6
7 8
```

## Usage
In order o try it out yourself you have to follow the following steps:

1. Clone the repo.
```bash
git clone git@github.com:roykn/tile-puzzle.git
```

2. Start the game and try to solve it.
```
go run tile-puzzle.go
```

3. Test the source.
```
go test -v
```
```
go test -cover
```
```
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```

## Contributing
Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License
Distributed under the MIT License. See `LICENSE` for more information.