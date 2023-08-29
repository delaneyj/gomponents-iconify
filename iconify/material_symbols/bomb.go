package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bomb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.65 22.8q-3.125 0-5.313-2.212T1.15 15.25q0-3.125 2.163-5.288T8.6 7.8h.325L9.6 6.625q.3-.55.9-.712t1.15.162l.75.425l.125-.2q.575-1.075 1.8-1.4t2.3.3l.875.5l-1 1.725l-.875-.5q-.35-.2-.763-.088t-.612.463l-.125.2l1 .575q.525.3.688.9t-.138 1.125L15 11.3q.575.9.863 1.913t.287 2.087q0 3.125-2.187 5.313T8.65 22.8ZM20 8.8v-2h3v2h-3Zm-5.5-5.5v-3h2v3h-2Zm4.875 2.025l-1.4-1.4L20.1 1.8l1.4 1.4l-2.125 2.125Z"/>`),
		g.Group(children),
	)
}