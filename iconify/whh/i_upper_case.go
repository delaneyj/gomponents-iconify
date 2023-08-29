package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M320 128h-63v768h63q27 0 45.5 19t18.5 45.5t-18.5 45T320 1024H64q-26 0-45-18.5t-19-45T19 915t45-19h65V128H64q-26 0-45-18.5t-19-45T19 19T64 0h256q27 0 45.5 19T384 64.5t-18.5 45T320 128z"/>`),
		g.Group(children),
	)
}