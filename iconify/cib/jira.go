package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31.099 15.104L16 0L.896 15.104a1.273 1.273 0 0 0 0 1.797L16 32l15.099-15.099c.5-.49.5-1.302 0-1.797zM16 20.734L11.266 16L16 11.271L20.729 16z"/>`),
		g.Group(children),
	)
}