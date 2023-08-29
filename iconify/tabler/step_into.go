package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepInto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v12m4-4l-4 4m-4-4l4 4m-1 5a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/>`),
		g.Group(children),
	)
}