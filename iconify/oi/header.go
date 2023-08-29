package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Header(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h.5c.28 0 .5.22.5.5v4c0 .28-.22.5-.5.5H0v1h3V6h-.5c-.28 0-.5-.22-.5-.5V4h3v1.5c0 .28-.22.5-.5.5H4v1h3V6h-.5c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5H7V0H4v1h.5c.28 0 .5.22.5.5V3H2V1.5c0-.28.22-.5.5-.5H3V0H0z"/>`),
		g.Group(children),
	)
}