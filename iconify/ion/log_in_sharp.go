package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogInSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 80H192a16 16 0 0 0-16 16v144h153.37l-64-64L288 153.37l91.31 91.32a16 16 0 0 1 0 22.62L288 358.63L265.37 336l64-64H176v144a16 16 0 0 0 16 16h240a16 16 0 0 0 16-16V96a16 16 0 0 0-16-16ZM64 240h112v32H64z"/>`),
		g.Group(children),
	)
}