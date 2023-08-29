package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sheriff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M831.5 384q-26.5 0-44.5-19L640 512l147 147q18-19 44.5-19t45.5 19t19 45t-19 45t-45.5 19t-45-19t-18.5-45H602L481 905q31 19 31 55q0 27-19 45.5t-45.5 18.5t-45-18.5T384 960q0-36 31-55L294 704H128q0 26-18.5 45t-45 19T19 749T0 703.5t19-45T64.5 640t44.5 19l147-147l-147-147q-18 19-44.5 19T19 365T0 319.5t19-45T64.5 256t45 18.5T128 320h166l121-201q-31-19-31-55q0-26 18.5-45t45-19T493 19t19 45q0 36-31 55l121 201h166q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5t-19 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}