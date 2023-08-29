package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M9.85786 38.1422C17.6684 45.9527 30.3316 45.9527 38.1421 38.1422L9.85786 9.85791C2.04738 17.6684 2.04738 30.3317 9.85786 38.1422Z"/><path d="M33.8994 33.8995L39.1698 11.9138"/><path d="M24 23.9999L37.1716 10.8284"/><path d="M14.1006 14.1005L36.0863 8.83008"/><path fill="#2F88FF" d="M44 8C44 10.2091 42.2091 12 40 12C39.7154 12 39.4377 11.9703 39.1699 11.9137C38.3984 11.7509 37.7089 11.3658 37.1716 10.8284C36.6342 10.2911 36.2491 9.60161 36.0863 8.83013C36.0297 8.56232 36 8.28463 36 8C36 5.79086 37.7909 4 40 4C42.2091 4 44 5.79086 44 8Z"/></g>`),
		g.Group(children),
	)
}