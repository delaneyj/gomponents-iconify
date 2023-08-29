package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feStar0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStar1" fill="currentColor" fill-rule="nonzero"><path id="feStar2" d="M12.5 17.925L6.629 21l1.121-6.512L3 9.875l6.564-.95L12.5 3l2.936 5.925l6.564.95l-4.75 4.613L18.371 21z"/></g></g>`),
		g.Group(children),
	)
}