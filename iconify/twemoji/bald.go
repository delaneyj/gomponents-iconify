package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bald(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FFDC5D" d="M34.896 36C30.618 13.677 16.169 2.725 0 1.195V36h34.896z"/>`),
		g.Group(children),
	)
}