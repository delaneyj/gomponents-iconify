package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fatarrowup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 383v512q0 53-37.5 90.5T896 1023H128q-53 0-90.5-37.5T0 895V383h1q0-41 39-70L418 29q39-29 94.5-29T607 29l378 284q39 29 39 70z"/>`),
		g.Group(children),
	)
}