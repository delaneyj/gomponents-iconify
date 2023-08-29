package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Media(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M244.4 69.8L174.5 0h-58.2l69.8 69.8h58.3zm151.2 0L325.8 0h-58.2l69.8 69.8h58.2zM418.9 0l69.8 69.8H512V0h-93.1zm0 162.9h-93.1l69.8-69.8h-58.2l-69.8 69.8h-93.1l69.8-69.8h-58.2l-69.8 69.8h-93l69.8-69.8H0v372.4C0 491.1 20.9 512 46.5 512h418.9c25.7 0 46.5-20.9 46.5-46.5V93.1h-23.3l-69.7 69.8zM23.3 0H0v69.8h93.1L23.3 0z"/>`),
		g.Group(children),
	)
}