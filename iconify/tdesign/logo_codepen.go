package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoCodepen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .807l11 7.15v8.086l-11 7.15l-11-7.15V7.957l11-7.15Zm-9 9.614v3.158L5.256 12L3 10.42Zm.788 5.048L11 20.157V16.02l-4-2.8l-3.212 2.25ZM8.744 12L12 14.28L15.256 12L12 9.72L8.744 12ZM13 7.98l4 2.8l3.212-2.248L13 3.842V7.98Zm-2-4.137L3.788 8.531L7 10.779l4-2.8V3.842Zm10 6.579L18.744 12L21 13.58v-3.16Zm-.788 5.048L17 13.221l-4 2.8v4.136l7.212-4.688Z"/>`),
		g.Group(children),
	)
}