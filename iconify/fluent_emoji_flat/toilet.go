package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toilet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#B4ACBC" d="M3.5 14A1.5 1.5 0 0 0 2 15.5v1a.5.5 0 0 0 .5.5h1L9 24l-2 4s-1 2 1 2h14c1 0 2-1 2-2v-8c0-1.5 2-3 2-3h-2.5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5h-20Z"/><path fill="#CDC4D6" d="M21 2a1 1 0 0 0-1 1v12c0 .768.289 1.47.764 2H3.2c.927 4.564 4.962 8 9.8 8c4.49 0 8.289-2.959 9.553-7.033c.146.022.295.033.447.033h2a3 3 0 0 0 3-3V3a1 1 0 0 0-1-1h-6Z"/></g>`),
		g.Group(children),
	)
}