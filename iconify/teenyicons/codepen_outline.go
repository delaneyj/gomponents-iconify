package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodepenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.29-.407a.5.5 0 0 0-.58 0L7.5.5Zm7 5h.5a.5.5 0 0 0-.21-.407l-.29.407Zm0 4l.29.407A.5.5 0 0 0 15 9.5h-.5Zm-7 5l-.29.407a.5.5 0 0 0 .58 0L7.5 14.5Zm-7-5H0a.5.5 0 0 0 .21.407L.5 9.5Zm0-4l-.29-.407A.5.5 0 0 0 0 5.5h.5ZM7.21.907l7 5l.58-.814l-7-5l-.58.814ZM14 5.5v4h1v-4h-1Zm.21 3.593l-7 5l.58.814l7-5l-.58-.814Zm-6.42 5l-7-5l-.58.814l7 5l.58-.814ZM1 9.5v-4H0v4h1ZM.79 5.907l7-5l-.58-.814l-7 5l.58.814Zm0 4l7-5l-.58-.814l-7 5l.58.814Zm6.42-5l7 5l.58-.814l-7-5l-.58.814Zm-7 1l7 5l.58-.814l-7-5l-.58.814Zm7.58 5l7-5l-.58-.814l-7 5l.58.814ZM7 .5v4h1v-4H7Zm0 10v4h1v-4H7Z"/>`),
		g.Group(children),
	)
}