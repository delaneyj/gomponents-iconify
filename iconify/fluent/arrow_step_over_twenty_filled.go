package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepOverTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.149 3.146a.5.5 0 0 0 0 .707L15.294 7H10c-2.932 0-5.593 1.64-6.936 4.043a.5.5 0 1 0 .873.488C5.106 9.439 7.436 8 10 8h5.293l-3.144 3.145a.5.5 0 1 0 .707.707l3.984-3.985a.499.499 0 0 0 .014-.721l-3.998-4a.5.5 0 0 0-.707 0ZM8 15a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/>`),
		g.Group(children),
	)
}