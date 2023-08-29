package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quotedown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 768H704q-53 0-90.5-37.5T576 640V448q0-53 37.5-90.5T704 320h128V192q0-26-18.5-45t-45-19t-45.5-18.5t-19-45T723 19t45-19h128q31 0 61 32t48.5 77t18.5 83v448q0 53-37.5 90.5T896 768zm-576 0H128q-53 0-90.5-37.5T0 640V448q0-53 37.5-90.5T128 320h128V192q0-26-19-45t-45.5-19t-45-18.5t-18.5-45T146.5 19T192 0h128q30 0 60.5 32t49 77t18.5 83v448q0 53-37.5 90.5T320 768z"/>`),
		g.Group(children),
	)
}