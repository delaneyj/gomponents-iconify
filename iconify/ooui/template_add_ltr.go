package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemplateAddLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 5V1h-2v4h-4v2h4v4h2V7h4V5z"/><path fill="currentColor" d="M0 17V5h8v2H2v8h12v-2h2v4z"/>`),
		g.Group(children),
	)
}