package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaScissorsFill0"><g id="evaScissorsFill1"><g id="evaScissorsFill2" fill="currentColor"><path d="M20.21 5.71a1 1 0 1 0-1.42-1.42l-6.28 6.31l-3.3-3.31A3 3 0 0 0 9.5 6a3 3 0 1 0-3 3a3 3 0 0 0 1.29-.3L11.1 12l-3.29 3.3A3 3 0 0 0 6.5 15a3 3 0 1 0 3 3a3 3 0 0 0-.29-1.26ZM6.5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><path d="M15.21 13.29a1 1 0 0 0-1.42 1.42l5 5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g></g>`),
		g.Group(children),
	)
}