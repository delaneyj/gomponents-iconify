package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 24a8 8 0 0 0-8 8v8H72v-8a8 8 0 0 0-16 0v192a8 8 0 0 0 16 0v-8h112v8a8 8 0 0 0 16 0V32a8 8 0 0 0-8-8Zm-16 160H80a8 8 0 0 1 0-16h96a8 8 0 0 1 0 16Zm0-48H80a8 8 0 0 1 0-16h96a8 8 0 0 1 0 16Zm0-48H80a8 8 0 0 1 0-16h96a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}