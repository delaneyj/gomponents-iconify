package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fontrounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736.992 192h-224v736q0 40-28 68t-68 28t-68-28t-28-68V192h-224q-40 0-68-28t-28-68t28-68t68-28h640q40 0 68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}