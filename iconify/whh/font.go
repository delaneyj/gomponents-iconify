package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.56 320q-26.5 0-45.5-19t-19-45q0-53-37.5-90.5t-90.5-37.5h-64q-26 0-45 18.5t-19 45.5v640q0 26 19 45t45.5 19t45 18.5t18.5 45t-18.5 45.5t-45.5 19h-384q-26 0-45-18.5t-19-45.5t19-45.5t45.5-18.5t45-19t18.5-45V192q0-27-18.5-45.5t-45.5-18.5h-64q-53 0-90.5 37.5t-37.5 90.5q0 26-18.5 45t-45 19t-45.5-19t-19-45V64q0-26 19-45t45-19h896q26 0 45 19t19 45v192q0 27-18.5 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}