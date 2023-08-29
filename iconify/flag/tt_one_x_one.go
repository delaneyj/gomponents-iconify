package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TtOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" d="M0 0h512v512H0z"/><g fill-rule="evenodd"><path fill="#e00000" d="M371 512L0 1v510.7l371 .3zM141 0l371 511V.2L141 0z"/><path d="M22.2.2h94.9l374.5 511.3h-97.9L22.2.2z"/></g>`),
		g.Group(children),
	)
}