package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearMeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.05 13.95l-6.475-2.625q-.325-.125-.475-.387t-.15-.538q0-.275.163-.537t.487-.388l15.35-5.7q.3-.125.575-.05T20 4q.2.2.275.475t-.05.575l-5.7 15.35q-.125.325-.388.488t-.537.162q-.275 0-.537-.15t-.388-.475L10.05 13.95Zm3.5 3.35L17.6 6.4L6.7 10.45l4.9 1.95l1.95 4.9Zm-1.95-4.9Z"/>`),
		g.Group(children),
	)
}