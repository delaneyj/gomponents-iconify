package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469.3 77H42.7C19.1 77 0 96.1 0 119.7v298.7C0 441.9 19.1 461 42.7 461h426.7c23.6 0 42.7-19.1 42.7-42.7V119.7C512 96.1 492.9 77 469.3 77zM384 354.3H128c-47.1 0-85.3-38.2-85.3-85.3s38.2-85.3 85.3-85.3s85.3 38.2 85.3 85.3c0 15.6-4.5 30.1-11.8 42.7h109c-7.3-12.6-11.8-27-11.8-42.7c0-47.1 38.2-85.3 85.3-85.3s85.3 38.2 85.3 85.3s-38.2 85.3-85.3 85.3z"/>`),
		g.Group(children),
	)
}