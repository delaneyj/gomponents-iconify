package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultipleForwardLeftBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18 10.13c-1.613-1.153-3.455-1.56-5-1.56V6.227c0-.958 0-1.437-.295-1.57c-.295-.132-.653.186-1.37.823L7.363 9.01C5.795 10.405 5.01 11.102 5.01 12c0 .899.785 1.596 2.353 2.99l3.973 3.53c.716.637 1.074.956 1.369.823c.295-.133.295-.611.295-1.57v-2.344c3.6 0 7.5 1.714 9 4.571c0-2.431-.377-4.378-1-5.925"/><path d="M8.461 4.5L3.245 9.344a3.897 3.897 0 0 0 .126 5.823L8.46 19.5"/></g>`),
		g.Group(children),
	)
}