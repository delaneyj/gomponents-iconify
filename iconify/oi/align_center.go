package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h8V0H0zm1 2v1h6V2H1zM0 4v1h8V4H0zm1 2v1h6V6H1z"/>`),
		g.Group(children),
	)
}