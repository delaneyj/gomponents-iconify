package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.853 19.147L12.707 12l7.146-7.147a.5.5 0 0 0-.707-.707L12 11.293L4.853 4.147a.5.5 0 0 0-.707.707L11.293 12l-7.147 7.146a.5.5 0 1 0 .707.707L12 12.707l7.146 7.147a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707z"/>`),
		g.Group(children),
	)
}