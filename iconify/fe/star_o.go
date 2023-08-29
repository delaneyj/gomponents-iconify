package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feStarO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feStarO1" fill="currentColor" fill-rule="nonzero"><path id="feStarO2" d="m9.282 17.362l3.218-1.685l3.218 1.685l-.615-3.57l2.604-2.527l-3.598-.52L12.5 7.496l-1.609 3.247l-3.598.52l2.604 2.529l-.615 3.57Zm3.218.563L6.629 21l1.121-6.512L3 9.875l6.564-.95L12.5 3l2.936 5.925l6.564.95l-4.75 4.613L18.371 21L12.5 17.925Z"/></g></g>`),
		g.Group(children),
	)
}