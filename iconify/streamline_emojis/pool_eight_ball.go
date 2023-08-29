package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoolEightBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M10.47 45.24a13.53 1.76 0 1 0 27.06 0a13.53 1.76 0 1 0-27.06 0Z" opacity=".15"/><path fill="#656769" d="M4 24a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#525252" d="M24 39.66A20 20 0 0 1 4.12 21.83C4.05 22.54 4 23.27 4 24a20 20 0 0 0 40 0c0-.73 0-1.46-.12-2.17A20 20 0 0 1 24 39.66Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M4 24a20 20 0 1 0 40 0a20 20 0 1 0-40 0Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.18 24a8.82 8.82 0 1 0 17.64 0a8.82 8.82 0 1 0-17.64 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.5 21.57a2.5 1.91 0 1 0 5 0a2.5 1.91 0 1 0-5 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M20.94 25.91a3.06 2.43 0 1 0 6.12 0a3.06 2.43 0 1 0-6.12 0Z"/><path fill="#87898c" d="M29 7.63c0 1.08-2.24 2-5 2s-5-.88-5-2s2.24-2 5-2s5 .91 5 2Z"/>`),
		g.Group(children),
	)
}