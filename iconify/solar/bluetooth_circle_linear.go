package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M14.2 9.593L11 12V7.623c0-.954 0-1.432.302-1.583c.301-.151.682.135 1.444.708L14.2 7.842c.533.401.8.602.8.876c0 .273-.267.474-.8.875Zm0 6.565l-1.454 1.094c-.762.573-1.143.86-1.444.708C11 17.809 11 17.331 11 16.377V12l3.2 2.407c.533.401.8.602.8.875c0 .274-.267.475-.8.876Z"/><path fill="currentColor" d="M8.48 8.924a.75.75 0 1 0-.96 1.152l.96-1.152Zm3 2.5l-3-2.5l-.96 1.152l3 2.5l.96-1.152Z"/><path fill="currentColor" d="M8.48 15.076a.75.75 0 0 1-.96-1.152l.96 1.152Zm3-2.5l-3 2.5l-.96-1.152l3-2.5l.96 1.152Z"/><circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="1.5"/></g>`),
		g.Group(children),
	)
}