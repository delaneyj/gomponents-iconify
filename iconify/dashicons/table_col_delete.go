package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableColDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.4 9.98L7.68 8.7v-.256L6.4 7.164V9.98zm6.4-1.532l1.28-1.28V9.92L12.8 8.64v-.192zm7.68 9.472V0H0v17.92h20.48zm-1.28-2.56h-5.12v-1.024l-.256.256l-1.024-1.024v1.792H7.68v-1.792l-1.024 1.024l-.256-.256v1.024H1.28V1.28H6.4v2.368l.704-.704l.576.576V1.216h5.12V3.52l.96-.96l.32.32V1.216h5.12V15.36zm-5.76-2.112l-3.136-3.136l-3.264 3.264l-1.536-1.536l3.264-3.264L5.632 5.44l1.536-1.536l3.136 3.136l3.2-3.2l1.536 1.536l-3.2 3.2l3.136 3.136l-1.536 1.536z"/>`),
		g.Group(children),
	)
}