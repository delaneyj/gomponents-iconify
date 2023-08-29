package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeypadFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaKeypadFill0"><g id="evaKeypadFill1"><path id="evaKeypadFill2" fill="currentColor" d="M5 2a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 6a3 3 0 1 0-3-3a3 3 0 0 0 3 3ZM5 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3ZM5 16a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm7 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3Z"/></g></g>`),
		g.Group(children),
	)
}