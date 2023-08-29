package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M5.5 1.5A1.5 1.5 0 0 1 7 0h2a1.5 1.5 0 0 1 1.5 1.5v11a1.5 1.5 0 0 1-1.404 1.497c.35.305.872.678 1.628 1.056A.5.5 0 0 1 10.5 16h-5a.5.5 0 0 1-.224-.947c.756-.378 1.277-.75 1.628-1.056A1.5 1.5 0 0 1 5.5 12.5v-11ZM7 1a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-11A.5.5 0 0 0 9 1H7Z"/><path d="M8.5 2.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm0 2a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm0 2a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm0 2a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/></g>`),
		g.Group(children),
	)
}