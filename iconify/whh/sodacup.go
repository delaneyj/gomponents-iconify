package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sodacup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m576 384l-64 576q-5 24-20 44t-36 20H184q-21 0-36.5-20T128 960L64 384q-26 0-45-18.5t-19-45T19 275t45-19h315l56-224q4-17 18-26t30-4.5t24 20t4 32.5l-50 202h115q26 0 45 19t19 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}