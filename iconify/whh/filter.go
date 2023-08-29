package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m992.06 128l-355 439l-61 393q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5l-61-393l-355-439l-23-32q-9-15-9-32q0-26 19-45t45-19h896q27 0 45.5 19t18.5 45q0 17-9 32z"/>`),
		g.Group(children),
	)
}