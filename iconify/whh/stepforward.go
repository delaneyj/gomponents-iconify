package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stepforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1009.86 551l-435 463q-25 27-62-13V24q37-39 62-12l435 463q15 16 15 38t-15 38zm-881 474q-53 0-90.5-37.5T.86 897V128q0-53 37.5-90.5T128.86 0t90.5 37.5t37.5 90.5v769q0 53-37.5 90.5t-90.5 37.5z"/>`),
		g.Group(children),
	)
}