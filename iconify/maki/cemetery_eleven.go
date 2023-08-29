package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CemeteryEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.65 8H8l1-5.61A.36.36 0 0 0 8.63 2H7.16c0-.65-.7-1-1.67-1s-1.83.35-1.83 1H2.35a.36.36 0 0 0-.35.39L3 8h-.65a.35.35 0 0 0-.35.36V10h7V8.36A.35.35 0 0 0 8.66 8h-.01zM7 5H4V4h3v1z" fill="currentColor"/>`),
		g.Group(children),
	)
}