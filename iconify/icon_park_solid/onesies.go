package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onesies(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOnesies0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M7 21L4 11c1.833-1.333 9-6 20-6s18 4.667 20 6l-3 10h-7v16c0 4-3 7-7 7h-6c-4 0-7-3-7-7V21H7Z"/><path stroke="#000" d="M34 34s-9 0-9 10M14 34s9 0 9 10m6.811-38.5a6 6 0 1 1-11.622 0"/><path stroke="#fff" d="M36 6.99C32.81 5.883 28.784 5 24 5s-8.843.883-12.054 1.99M14 31v6m20-6v6m-13 7h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOnesies0)"/>`),
		g.Group(children),
	)
}