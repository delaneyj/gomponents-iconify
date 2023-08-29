package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsTwoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPills2Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM26 16c0 5.523-4.477 10-10 10S6 21.523 6 16S10.477 6 16 6s10 4.477 10 10Zm-5.153-1.997a1 1 0 0 1-.543 1.305l-7.845 3.233a1 1 0 1 1-.763-1.849l7.846-3.233a1 1 0 0 1 1.305.544ZM42 32c0 5.523-4.477 10-10 10s-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10Zm-4.89 1.175a1 1 0 0 1-1.2.75l-8.269-1.901a1 1 0 1 1 .448-1.95l8.27 1.902a1 1 0 0 1 .75 1.199Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPills2Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}