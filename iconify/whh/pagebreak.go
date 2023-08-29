package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pagebreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m816.417 400l-112 112h-32l-112-112l-112 112h-32l-112-112l-112 112h-32l-112-112l-48 48V64q0-26 19-45t45-19h768q27 0 45.5 19t18.5 45v416zm-784 112h32l112 112l112-112h32l112 112l112-112h32l112 112l112-112h32l64 64v384q0 26-18.5 45t-45.5 19h-768q-26 0-45-19t-19-45V544z"/>`),
		g.Group(children),
	)
}