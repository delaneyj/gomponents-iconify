package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M84.437 39.721H60.273V15.563a1.814 1.814 0 0 0-1.812-1.813H41.536a1.813 1.813 0 0 0-1.812 1.813l-.001 24.16l-24.159-.001c-.961 0-1.812.851-1.813 1.813V58.46a1.81 1.81 0 0 0 1.813 1.812h24.16v24.165a1.814 1.814 0 0 0 1.813 1.813H58.46a1.813 1.813 0 0 0 1.813-1.813V60.273l24.163-.001a1.81 1.81 0 0 0 1.813-1.813l.001-16.925a1.813 1.813 0 0 0-1.813-1.813z"/>`),
		g.Group(children),
	)
}