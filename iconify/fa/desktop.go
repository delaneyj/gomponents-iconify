package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1792 992V160q0-13-9.5-22.5T1760 128H160q-13 0-22.5 9.5T128 160v832q0 13 9.5 22.5t22.5 9.5h1600q13 0 22.5-9.5t9.5-22.5zm128-832v1088q0 66-47 113t-113 47h-544q0 37 16 77.5t32 71t16 43.5q0 26-19 45t-45 19H704q-26 0-45-19t-19-45q0-14 16-44t32-70t16-78H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1600q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}