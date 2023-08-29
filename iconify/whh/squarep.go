package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-640q0-53-37.5-90.5t-90.5-37.5h-224q-13 0-22.5 9.5t-9.5 22.5v448q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V576h192q53 0 90.5-37.5t37.5-90.5v-64zm-128 128h-192V320h192q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}