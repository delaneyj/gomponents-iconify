package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainerSoftware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 12h-8V4h8Zm-6-2h4V6h-4Z"/><path fill="currentColor" d="M17 15V9H9v14h14v-8Zm-6-4h4v4h-4Zm4 10h-4v-4h4Zm6 0h-4v-4h4Z"/><path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h10v2H6v20h20V16h2v10a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}