package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crowdfire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2v2h8v8h2c0 5.523-4.477 10-10 10S2 17.523 2 12Zm16 0h-2V8h-4V6a6 6 0 1 0 6 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}