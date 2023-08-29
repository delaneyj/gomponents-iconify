package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h7v6l-4 7M20 8h-7v6l4 7M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-1 3h2v5h-2V8Z"/>`),
		g.Group(children),
	)
}