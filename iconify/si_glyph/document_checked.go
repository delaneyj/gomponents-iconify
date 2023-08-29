package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.024.016v3.979h4.947L9.024.016z"/><path d="M8 5.062V0H2.008v12.745l.436.437l2.135-2.136l2.109 2.105l-2.801 2.802H14V5l-6 .062z"/><path d="m4.58 12.46l-2.136 2.136l-1.399-1.4l-.693.693l2.092 2.092l2.829-2.829l-.693-.692z"/></g>`),
		g.Group(children),
	)
}