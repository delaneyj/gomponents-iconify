package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsThermometerNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18.982 28.204v1.78a8 8 0 1 0 10-.008v-1.78l-.006-6.211l-.01-12.408c-.002-2.958-2.243-5.355-5.004-5.353c-2.762.003-4.998 2.403-4.996 5.362l.01 12.407l.006 6.212Zm7.995-5.921c-1.457.025-2.823-.328-4.018-.637c-.73-.188-1.397-.36-1.983-.42l-.01-11.642c-.001-1.775 1.34-3.215 2.998-3.217c1.657-.001 3 1.437 3.002 3.212v.619l-2 .002a1 1 0 0 0 .002 2l2-.002l.002 2l-2 .002a1 1 0 0 0 .002 2l2-.002l.001 2l-2 .002a1 1 0 0 0 .002 2l2-.002l.002 2.085Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsThermometerNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}