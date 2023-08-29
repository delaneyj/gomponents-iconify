package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M64 57.1a6.899 6.899 0 0 1-6.898 6.903H6.892A6.9 6.9 0 0 1-.004 57.1V6.9A6.899 6.899 0 0 1 6.892.001h50.21A6.898 6.898 0 0 1 64 6.9v50.2"/><path fill="#fff" d="m47.972 23.11l2.205 24.353l-24.339-1.267c-1.715-1.932-1.975-4.567-.521-6.082l2.467-2.569l-12.07-11.578a5.103 5.103 0 0 1-.152-7.208l4.58-4.768a5.101 5.101 0 0 1 7.208-.15L39.41 25.42l2.463-2.573c1.453-1.511 4.093-1.365 6.094.263"/>`),
		g.Group(children),
	)
}