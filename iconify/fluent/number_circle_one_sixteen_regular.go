package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleOneSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5.5a.5.5 0 0 0-.99-.098c-.113.565-.27.85-.446 1.047c-.17.19-.371.324-.685.531l-.156.104a.5.5 0 0 0 .554.832l.142-.093c.193-.127.393-.259.581-.413v3.09a.5.5 0 1 0 1 0v-5ZM8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2ZM3 8a5 5 0 1 1 10 0A5 5 0 0 1 3 8Z"/>`),
		g.Group(children),
	)
}