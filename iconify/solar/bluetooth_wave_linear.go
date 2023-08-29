package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothWaveLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M13.263 8.262L8 12V5.225c0-1.887 0-2.83.605-3.14c.604-.309 1.376.24 2.92 1.336l1.738 1.234C14.421 5.477 15 5.888 15 6.46c0 .57-.579.981-1.737 1.803Zm0 11.083l-1.738 1.234c-1.544 1.096-2.316 1.645-2.92 1.335C8 21.605 8 20.662 8 18.775V12l5.263 3.738C14.421 16.56 15 16.97 15 17.54c0 .57-.579.982-1.737 1.804Z"/><path fill="currentColor" d="M2.57 14.886a.75.75 0 1 0 .86 1.228l-.86-1.228Zm.86 1.228l5-3.5l-.86-1.228l-5 3.5l.86 1.228Z"/><path fill="currentColor" d="M2.57 9.114a.75.75 0 0 1 .86-1.228l-.86 1.228Zm.86-1.228l5 3.5l-.86 1.228l-5-3.5l.86-1.228Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M19 5s2 2.1 2 7s-2 7-2 7M17 8s1 1.9 1 4s-1 4-1 4"/></g>`),
		g.Group(children),
	)
}