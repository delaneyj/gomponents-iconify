package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTextMessage0"><g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M44 7H4v30h15l5 5l5-5h15V7Z"/><path stroke="#000" d="M14 16h6m-6 8h2m13-10l7 14m-7-14l-7 14m2-4h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTextMessage0)"/>`),
		g.Group(children),
	)
}