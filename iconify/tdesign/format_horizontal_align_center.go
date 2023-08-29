package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHorizontalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2v20H4V2h2Zm14 0v20h-2V2h2Zm-7 3v14h-2V5h2Z"/>`),
		g.Group(children),
	)
}