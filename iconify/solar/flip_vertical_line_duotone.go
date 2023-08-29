package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5.886 2h12.228c1.702 0 2.553 0 2.832.542c.28.543-.216 1.235-1.205 2.62l-1.13 1.582c-.44.616-.66.924-.982 1.09C17.306 8 16.927 8 16.17 8H7.83c-.757 0-1.136 0-1.459-.166c-.323-.166-.543-.474-.983-1.09l-1.13-1.582c-.988-1.385-1.483-2.077-1.204-2.62C3.334 2 4.184 2 5.886 2Zm0 20h12.228c1.702 0 2.553 0 2.832-.542c.28-.543-.216-1.235-1.205-2.62l-1.13-1.582c-.44-.616-.66-.924-.982-1.09C17.306 16 16.927 16 16.17 16H7.83c-.757 0-1.136 0-1.459.166c-.323.166-.543.474-.983 1.09l-1.13 1.581c-.988 1.386-1.483 2.078-1.204 2.62c.28.543 1.13.543 2.832.543Z"/><path stroke-linecap="round" d="M10 12h4m4 0h4M2 12h4" opacity=".5"/></g>`),
		g.Group(children),
	)
}