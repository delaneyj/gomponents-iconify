package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 5)"><path d="M15.031 2.016v1.015H1.938V.078H.009v5.893h1.929V4.938h13.093v1.033h.938V2.016h-.938z"/><ellipse cx="3.985" cy=".97" rx=".985" ry=".97"/><path d="M13.991 1.998H6.013v-.866c0-.598.643-1.083 1.434-1.083h5.109c.793 0 1.435.485 1.435 1.083v.866z"/></g>`),
		g.Group(children),
	)
}