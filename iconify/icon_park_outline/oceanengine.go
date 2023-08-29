package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oceanengine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="m21.483 8.474l2.397 8.153H4.758l2.458-8.153h14.267Zm-12.35 9.651l8.213 1.979l-9.531 16.544L2 30.474l7.133-12.349Zm2.159 15.526l5.815-6.174l9.591 16.545L18.425 46l-7.133-12.349Zm28.773 5.815H25.798L23.4 31.373h19.123m2.757-13.907l-7.133 12.409l-8.212-1.979l9.53-16.545m-3.476 2.998l-5.815 6.174l-9.591-16.605L28.856 2l7.133 12.349Z"/>`),
		g.Group(children),
	)
}