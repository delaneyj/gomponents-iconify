package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.986 12.806V23.2l-5.197-5.197zM6.796 6.01H17.19l-5.197 5.197zm9.275-1.12H5.677L.8.015h10.394zm7.116 7.117L17.99 6.81l-5.197 5.197l5.197 5.197zm.813-.813L12.806 0H24zM0 23.2V.813l11.194 11.194zm23.187.8H.8l11.193-11.194Z"/>`),
		g.Group(children),
	)
}