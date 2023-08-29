package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisappointedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDisappointedFace0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="m33 25l-4-2m-11 0l-4 2m17 10s-2-4-7-4s-7 4-7 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDisappointedFace0)"/>`),
		g.Group(children),
	)
}