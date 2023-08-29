package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumberedRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 15h11v2H2zm0-6h11v2H2zm0-6h11v2H2zm15-2h-1v1h1v4h1V1zm-2 12v1h2v1h-1v1h1v1h-2v1h3v-5zm0-6v1h2v1h-1.5c-.3 0-.5.2-.5.5V12h3v-1h-2v-1h1.5c.3 0 .5-.2.5-.5v-2c0-.3-.2-.5-.5-.5z"/>`),
		g.Group(children),
	)
}