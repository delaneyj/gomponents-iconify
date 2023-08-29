package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm8.5-2.5a1.5 1.5 0 1 1 2.56 1.06l-3.828 3.83a2.5 2.5 0 0 0-.732 1.767V17a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-3.975a.5.5 0 0 1 .121-.197l3.829-3.828A3.5 3.5 0 1 0 8.5 9.5a1 1 0 1 0 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}