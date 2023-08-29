package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Legacyfilemanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.38 896h-768q-53 0-90.5-37.5T.38 768V256q0-26 18.5-45t45.5-19h480l46-85q7-18 28-30.5t44-12.5h240q24 0 45 12.5t28 30.5l49 85v576q0 53-37.5 90.5t-90.5 37.5zm0-640q0-26-19-45t-45-19h-128q-27 0-45.5 19t-18.5 45.5t-28 45t-68 18.5h-384q-13 0-22.5 9.5t-9.5 22.5v384q0 13 9.5 22.5t22.5 9.5h704q13 0 22.5-9.5t9.5-22.5V256zm-850-212q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84H.38z"/>`),
		g.Group(children),
	)
}