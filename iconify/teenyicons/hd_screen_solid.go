package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdScreenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.948 1.108A.744.744 0 0 0 0 1.823v9.354c0 .494.473.85.948.715A23.85 23.85 0 0 1 7 10.98V13H2v1h11v-1H8v-2.02c2.039.042 4.073.346 6.052.912a.744.744 0 0 0 .948-.715V1.823a.744.744 0 0 0-.948-.715a23.85 23.85 0 0 1-13.104 0Z"/>`),
		g.Group(children),
	)
}