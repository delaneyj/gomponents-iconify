package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryMediumLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 58H32a22 22 0 0 0-22 22v96a22 22 0 0 0 22 22h168a22 22 0 0 0 22-22V80a22 22 0 0 0-22-22Zm10 118a10 10 0 0 1-10 10H32a10 10 0 0 1-10-10V80a10 10 0 0 1 10-10h168a10 10 0 0 1 10 10ZM102 96v64a6 6 0 0 1-12 0V96a6 6 0 0 1 12 0Zm-40 0v64a6 6 0 0 1-12 0V96a6 6 0 0 1 12 0Zm192 0v64a6 6 0 0 1-12 0V96a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}