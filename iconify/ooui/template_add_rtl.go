package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemplateAddRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 7h4v4h2V7h4V5H6V1H4v4H0z"/><path fill="currentColor" d="M4 13h2v2h12V7h-6V5h8v12H4z"/>`),
		g.Group(children),
	)
}