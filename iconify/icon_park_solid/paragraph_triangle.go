package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSParagraphTriangle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 9h18M24 19h18M6 29h36M6 39h36"/><path fill="#fff" d="M6 9.766a1 1 0 0 1 1.514-.857l7.057 4.233a1 1 0 0 1 0 1.716L7.515 19.09A1 1 0 0 1 6 18.234V9.766Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSParagraphTriangle0)"/>`),
		g.Group(children),
	)
}