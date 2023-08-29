package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighVoltage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M10.9 45.29a13.1 1.71 0 1 0 26.2 0a13.1 1.71 0 1 0-26.2 0Z" opacity=".15"/><path fill="#ffe500" d="M33.72 18.24H26l1.71-14a.57.57 0 0 0-1-.39l-13.32 19a1.14 1.14 0 0 0 .93 1.79H22l-1.69 13.72a.57.57 0 0 0 1 .4L34.65 20a1.14 1.14 0 0 0-.93-1.76Z"/><path fill="#fff48c" d="M26.56 22h6.69l1.4-2a1.14 1.14 0 0 0-.93-1.8H27Zm-11.25 2.59l11.93-17l.06-.07l.4-3.31a.57.57 0 0 0-1-.39l-13.32 19a1.14 1.14 0 0 0 .93 1.79Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.72 18.24H26l1.71-14a.57.57 0 0 0-1-.39l-13.32 19a1.14 1.14 0 0 0 .93 1.79H22l-1.69 13.72a.57.57 0 0 0 1 .4L34.65 20a1.14 1.14 0 0 0-.93-1.76Z"/>`),
		g.Group(children),
	)
}