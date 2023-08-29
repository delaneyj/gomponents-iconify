package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSThin0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path fill="#fff" d="M33 25H15c-1.59 0-3.77-.23-5-1c-3.65-2.31-4.34-7.37-2-11l6-9c3.93 8.43 16.04 8.42 20 0l6 9c2.34 3.63 1.64 8.69-2 11c-1.23.78-3.41 1-5 1Z"/><path d="m12 25l2.52 9.55c.87 3.38-.06 6.97-2.52 9.45m24-19l-2.52 9.58c-.87 3.38.06 6.94 2.52 9.42m-12-2v-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSThin0)"/>`),
		g.Group(children),
	)
}