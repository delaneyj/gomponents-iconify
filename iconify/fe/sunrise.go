package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSunrise0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSunrise1" fill="currentColor" fill-rule="nonzero"><path id="feSunrise2" d="M21 15H3a1 1 0 0 1 0-2h1a8 8 0 1 1 16 0h1a1 1 0 0 1 0 2Zm-3 4H6a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2ZM6 13h12a6 6 0 1 0-12 0Z"/></g></g>`),
		g.Group(children),
	)
}