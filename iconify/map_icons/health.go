package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Health(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M49 17c0-.55-.45-1-1-1H35c-.55 0-1-.45-1-1V2c0-.55-.45-1-1-1H17c-.55 0-1 .45-1 1v13c0 .55-.45 1-1 1H2c-.55 0-1 .45-1 1v16c0 .55.45 1 1 1h13c.55 0 1 .45 1 1v13c0 .55.45 1 1 1h16c.55 0 1-.45 1-1V35c0-.55.45-1 1-1h13c.55 0 1-.45 1-1V17z"/>`),
		g.Group(children),
	)
}