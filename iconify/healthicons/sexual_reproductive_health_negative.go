package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SexualReproductiveHealthNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsSexualReproductiveHealthNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm17.546 15.48a7 7 0 1 0 4.372 5.454a1 1 0 0 1 1.977-.304A9 9 0 0 1 16 30.944V33h2a1 1 0 1 1 0 2h-2v4a1 1 0 1 1-2 0v-4h-2a1 1 0 1 1 0-2h2v-2c0-.019 0-.037.002-.055a9.004 9.004 0 0 1-6.024-3.316a9 9 0 0 1 10.295-14.013a1 1 0 1 1-.727 1.863Zm6.86.018a7 7 0 1 1-.177 12.93a1 1 0 1 0-.791 1.837a9 9 0 0 0 11.018-13.306l5.513-5.514l-.032 2.542a1 1 0 1 0 2 .026l.076-6.026l-6.026.076a1 1 0 1 0 .026 2l2.542-.032l-5.401 5.401a8.999 8.999 0 0 0-13.89 1.967a9 9 0 0 0-1.201 5.662a1 1 0 0 0 1.986-.236a7 7 0 0 1 4.357-7.327Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsSexualReproductiveHealthNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}