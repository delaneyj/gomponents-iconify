package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMessageSquareFill0"><g id="evaMessageSquareFill1"><path id="evaMessageSquareFill2" fill="currentColor" d="M19 3H5a3 3 0 0 0-3 3v15a1 1 0 0 0 .51.87A1 1 0 0 0 3 22a1 1 0 0 0 .51-.14L8 19.14a1 1 0 0 1 .55-.14H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM8 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm4 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`),
		g.Group(children),
	)
}