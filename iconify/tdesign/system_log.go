package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemLog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v22H1V1Zm2 8.667V21h18V9.667H3Zm18-2V3H3v4.667h18ZM5 4h2.004v2.004H5V4Zm1 8h12v2H6v-2Zm0 4h6v2H6v-2Z"/>`),
		g.Group(children),
	)
}