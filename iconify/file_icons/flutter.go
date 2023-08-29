package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M415.513 0L79.18 337.114L0 257.89L257.29 0h158.223m-.972 238.328H256.32L118.506 376.46L253.804 512h156.97L276.516 376.674l138.025-138.346z"/>`),
		g.Group(children),
	)
}