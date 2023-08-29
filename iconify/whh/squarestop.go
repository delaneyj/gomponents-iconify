package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarestop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-128-704q0-26-18.5-45t-45.5-19h-384q-26 0-45 19t-19 45v384q0 27 19 45.5t45 18.5h384q27 0 45.5-18.5t18.5-45.5V320z"/>`),
		g.Group(children),
	)
}