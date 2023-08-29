package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#D20000" d="M.5.5h300v150H.5z"/><path fill="#FFE600" d="M.5.5h45l105 72.321L255.5.5h45l-300 150h45l105-72.321l105 72.321h45L.5.5zm300 60v30l-300-30v30l300-30zm-165-60l15 64.286L165.5.5h-30zm0 150l15-64.286l15 64.286h-30z"/><circle cx="150.5" cy="75.5" r="24.107" fill="#FFE600" stroke="#D20000" stroke-width="4.286"/></g>`),
		g.Group(children),
	)
}