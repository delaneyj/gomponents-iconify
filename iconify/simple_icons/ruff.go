package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.683 11.593l-8.51-2.14l8.34-9.066a.23.23 0 0 0-.29-.352L2.252 11.62a.227.227 0 0 0-.108.226a.23.23 0 0 0 .164.19l8.497 2.497l-8.35 9.08a.228.228 0 0 0-.007.303a.227.227 0 0 0 .3.047l19-11.953a.228.228 0 0 0 .105-.23a.225.225 0 0 0-.172-.187z"/>`),
		g.Group(children),
	)
}