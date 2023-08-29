package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v20h-2V2h2ZM2 2h16v20H2V2Zm2 2v16h12V4H4Z"/>`),
		g.Group(children),
	)
}