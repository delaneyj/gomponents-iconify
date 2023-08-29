package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M37 1c-.55 0-1 .45-1 1v14c0 .55-.45 1-1 1H15c-.55 0-1-.45-1-1V2c0-.55-.45-1-1-1H4c-.55 0-1 .45-1 1v46c0 .55.45 1 1 1h9c.55 0 1-.45 1-1V32c0-.55.45-1 1-1h20c.55 0 1 .45 1 1v16c0 .55.45 1 1 1h9c.55 0 1-.45 1-1V2c0-.55-.45-1-1-1h-9z"/>`),
		g.Group(children),
	)
}