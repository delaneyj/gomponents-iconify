package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerplug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 827v133q0 27-19 45.5t-45.5 18.5t-45-18.5T384 960V827q-107-16-195-78T50.5 591T0 384q0-26 18.5-45T64 320h768q26 0 45 19t19 45q0 111-50.5 207T707 749.5T512 827zm128-571V64q0-26 18.5-45t45-19T749 19t19 45v192H640zm-512 0V64q0-26 18.5-45t45-19T237 19t19 45v192H128z"/>`),
		g.Group(children),
	)
}