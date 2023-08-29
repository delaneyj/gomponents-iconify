package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928.784 1024h-832q-40 0-68-28t-28-68V96q0-40 28-68t68-28t68 28t28 68v736h736q40 0 68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}