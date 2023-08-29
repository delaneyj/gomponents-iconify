package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulletpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 6.004H2V4h2.004v2.004H3Zm0 7H2V11h2.004v2.004H3Zm-1 7h2.004V18H2v2.004ZM8 4H7v2h15V4H8Zm-1 7h15v2H7v-2Zm1 7H7v2h15v-2H8Z"/>`),
		g.Group(children),
	)
}