package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.001 17a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Zm-13-15h4v20h-4V2Z"/>`),
		g.Group(children),
	)
}