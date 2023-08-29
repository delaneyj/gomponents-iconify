package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.65 22.8q-3.125 0-5.313-2.212T1.15 15.25q0-3.125 2.163-5.288T8.6 7.8h.325l1.425-2.475L12.4 6.5l.125-.2q.575-1.075 1.8-1.4t2.3.3l.875.5l-1 1.725l-.875-.5q-.35-.2-.763-.087t-.612.462l-.125.2L16.4 8.8L15 11.3q.575.9.862 1.913t.288 2.087q0 3.125-2.187 5.313T8.65 22.8Zm0-2q2.275 0 3.888-1.613T14.15 15.3q0-.775-.213-1.525T13.3 12.35l-.65-1.025l1.05-1.8l-2.6-1.5l-1.05 1.8h-1.1q-2.35 0-4.087 1.5T3.125 15.25q0 2.3 1.613 3.925T8.65 20.8ZM20 8.8v-2h3v2h-3Zm-5.5-5.5v-3h2v3h-2Zm4.875 2.025l-1.4-1.4L20.1 1.8l1.4 1.4l-2.125 2.125ZM8.65 15.3Z"/>`),
		g.Group(children),
	)
}