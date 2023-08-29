package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0C2.9 0 2 .9 2 2v1H1v4h6V3H6V2c0-1.1-.9-2-2-2zm0 1c.56 0 1 .44 1 1v1H3V2c0-.56.44-1 1-1z"/>`),
		g.Group(children),
	)
}