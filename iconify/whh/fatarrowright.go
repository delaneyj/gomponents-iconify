package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fatarrowright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m995.86 607l-284 378q-29 39-71 39l-512 1q-53 0-90.5-37.5T.86 897V128q0-53 37.5-90.5T128.86 0h512q42 0 71 39l284 378q29 40 29 95t-29 95z"/>`),
		g.Group(children),
	)
}