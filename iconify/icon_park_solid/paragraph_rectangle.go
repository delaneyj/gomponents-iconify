package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParagraphRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSParagraphRectangle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 9h18M24 19h18M6 29h36M6 39h36"/><path fill="#fff" d="M6 9h10v10H6z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSParagraphRectangle0)"/>`),
		g.Group(children),
	)
}