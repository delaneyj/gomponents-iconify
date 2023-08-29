package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviewOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPreviewOpen0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M24 36c11.046 0 20-12 20-12s-8.954-12-20-12S4 24 4 24s8.954 12 20 12Z"/><path d="M24 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPreviewOpen0)"/>`),
		g.Group(children),
	)
}