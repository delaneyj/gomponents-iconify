package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feathub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11.641.016h8.719v9.156h-8.719zm0 22.812h8.719v9.156h-8.719zM.016 11.641h31.969v8.719H.016z"/>`),
		g.Group(children),
	)
}