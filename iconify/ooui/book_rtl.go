package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 2v15h4a7.65 7.65 0 0 1 5 2a7.65 7.65 0 0 1 5-2h4V2h-4a7.65 7.65 0 0 0-5 2a7.65 7.65 0 0 0-5-2zm1.5 1.5H5C8 3.5 9 5 9 5v11.5a4.38 4.38 0 0 0-3-1H2.5z"/><path fill="currentColor" d="M9 3.5h2v1H9z"/>`),
		g.Group(children),
	)
}