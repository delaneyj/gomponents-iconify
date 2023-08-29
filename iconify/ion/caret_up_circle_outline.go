package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m342.43 273.77l-74.13-89.09a16 16 0 0 0-24.6 0l-74.13 89.09A16 16 0 0 0 181.86 300h148.28a16 16 0 0 0 12.29-26.23Z"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="M448 256c0-106-86-192-192-192S64 150 64 256s86 192 192 192s192-86 192-192Z"/>`),
		g.Group(children),
	)
}