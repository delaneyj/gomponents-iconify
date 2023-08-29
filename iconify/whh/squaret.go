package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squaret(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-288-768h-256q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h96v416q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V320h96q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5z"/>`),
		g.Group(children),
	)
}