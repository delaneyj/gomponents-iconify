package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 0H3c-.55 0-1 .45-1 1v14c0 .55.45 1 1 1h9c.55 0 1-.45 1-1V1c0-.55-.45-1-1-1zM7.5 15.278a.778.778 0 1 1 0-1.555a.778.778 0 0 1 0 1.555zM12 13H3V2h9v11z"/>`),
		g.Group(children),
	)
}