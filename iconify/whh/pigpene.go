package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpene(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928.784 1024h-832q-40 0-68-28t-28-68V96q0-40 28-68t68-28h832q40 0 68 28t28 68v832q0 40-28 68t-68 28zm-96-832h-640v640h640V192z"/>`),
		g.Group(children),
	)
}