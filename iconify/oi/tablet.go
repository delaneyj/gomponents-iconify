package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.34 0A.35.35 0 0 0 0 .34v7.31c0 .18.16.34.34.34h6.31c.18 0 .34-.16.34-.34V.34A.35.35 0 0 0 6.65 0H.34zM1 1h5v5H1V1zm2.5 5.5c.38 0 .63.42.44.75s-.68.33-.88 0c-.19-.33.05-.75.44-.75z"/>`),
		g.Group(children),
	)
}