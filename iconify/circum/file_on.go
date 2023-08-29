package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.485 7.35l-4.97-4.86a1.466 1.466 0 0 0-1.05-.43h-6.9a2.5 2.5 0 0 0-2.5 2.5v14.88a2.507 2.507 0 0 0 2.5 2.5h10.87a2.507 2.507 0 0 0 2.5-2.5V8.42a1.49 1.49 0 0 0-.45-1.07Zm-1.27.15h-2.34a1.5 1.5 0 0 1-1.5-1.5V3.75Zm.72 11.94a1.5 1.5 0 0 1-1.5 1.5H6.565a1.5 1.5 0 0 1-1.5-1.5V4.56a1.5 1.5 0 0 1 1.5-1.5h6.81V6a2.5 2.5 0 0 0 2.5 2.5h3.06Z"/>`),
		g.Group(children),
	)
}