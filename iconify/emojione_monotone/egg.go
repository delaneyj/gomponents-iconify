package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.56 29.807C51.111 11.768 39.666 2 32 2c-7.668 0-19.112 9.768-21.559 27.807C7.988 47.842 15.853 62 32 62c16.146 0 24.013-14.158 21.56-32.193z"/>`),
		g.Group(children),
	)
}