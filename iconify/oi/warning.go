package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.09 0c-.06 0-.1.04-.13.09L.02 6.9c-.02.05-.03.13-.03.19v.81c0 .05.04.09.09.09h6.81c.05 0 .09-.04.09-.09v-.81c0-.05-.01-.14-.03-.19L4.01.09A.142.142 0 0 0 3.88 0h-.81zM3 3h1v2H3V3zm0 3h1v1H3V6z"/>`),
		g.Group(children),
	)
}