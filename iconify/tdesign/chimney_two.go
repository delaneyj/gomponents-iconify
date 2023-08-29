package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChimneyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2H4v5.924L1.834 22H22V10H11.32L11 7.924V2ZM9 4v4.076L9.296 10H8v10H4.165L6 8.076V4h3Zm1 16v-8h10v8h-4v-4h-2v4h-4Z"/>`),
		g.Group(children),
	)
}