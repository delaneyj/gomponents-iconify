package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v2h-1v4h3v14H2V8h3V4H4V2Zm3 2v6H4v10h5v-6h6v6h5V10h-3V4H7Zm6 16v-4h-2v4h2Zm0-15v2h2v2h-2v2h-2V9H9V7h2V5h2Z"/>`),
		g.Group(children),
	)
}