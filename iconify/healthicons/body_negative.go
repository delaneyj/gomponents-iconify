package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBodyNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm13.52 3.931a2 2 0 1 0-1.04-3.862c-4.978 1.339-8.751 1.946-12.472 1.93c-3.727-.015-7.509-.655-12.512-1.937a2 2 0 0 0-.992 3.876c2.73.699 5.175 1.23 7.496 1.58V42a2 2 0 0 0 3.992.181l1-11A2 2 0 0 0 23 31h2a2 2 0 0 0 .008.181l1 11A2 2 0 0 0 30 42V19.554c2.325-.344 4.779-.885 7.52-1.623Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBodyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}