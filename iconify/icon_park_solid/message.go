package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Message(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMessage0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M44 24c0 11.046-8.954 20-20 20H4V24C4 12.954 12.954 4 24 4s20 8.954 20 20Z"/><path stroke="#000" d="M14 18h18m-18 8h18m-18 8h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMessage0)"/>`),
		g.Group(children),
	)
}