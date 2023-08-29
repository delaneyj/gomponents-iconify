package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 1024H96q-40 0-68-28T0 928t28-68t68-28h736V96q0-40 28-68t68-28t68 28t28 68v832q0 40-28 68t-68 28z"/>`),
		g.Group(children),
	)
}