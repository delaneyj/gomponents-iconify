package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DubleCornerArrowFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325.8.1L256 69.9h139.6L256 209.5l46.5 46.5l139.6-139.6V256l69.8-69.8L512 0L325.8.1zM209.5 256L69.9 395.6V256L.1 325.8L0 512l186.2-.1l69.8-69.8H116.4L256 302.5L209.5 256z"/>`),
		g.Group(children),
	)
}