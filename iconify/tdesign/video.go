package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm4 2.37L17.75 12L8 17.63V6.37Zm2 3.465v4.33L13.75 12L10 9.835Z"/>`),
		g.Group(children),
	)
}