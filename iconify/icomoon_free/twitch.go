package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1.5 0L0 2.5V14h4v2h2l2-2h2.5L15 9.5V0H1.5zM13 8.5L10.5 11H8l-2 2v-2H3V2h10v6.5z"/><path fill="currentColor" d="M9.5 4H11v4H9.5V4zm-3 0H8v4H6.5V4z"/>`),
		g.Group(children),
	)
}