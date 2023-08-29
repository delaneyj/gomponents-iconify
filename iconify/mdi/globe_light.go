package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.1 10c1-1 2.4-1.7 3.9-1.9V2h2v6.1c1.5.2 2.9.9 3.9 1.9H7.1m-1.8 3c-.2.6-.3 1.3-.3 2c0 3.9 3.1 7 7 7s7-3.1 7-7c0-.7-.1-1.4-.3-2H5.3Z"/>`),
		g.Group(children),
	)
}