package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Location(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feLocation0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feLocation1" fill="currentColor" fill-rule="nonzero"><path id="feLocation2" d="M12 19c.437 0 1.479-1.187 2.411-3.312C15.357 13.534 16 10.874 16 9.353C16 6.924 14.183 5 12 5S8 6.924 8 9.353c0 1.52.643 4.181 1.589 6.335C10.52 17.813 11.563 19 12 19Zm0 2c-3.314 0-6-8.138-6-11.647C6 5.844 8.686 3 12 3s6 2.844 6 6.353S15.314 21 12 21Zm0-10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}