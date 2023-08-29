package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouthNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsMouthNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 14.182C15.2 10.254 7 18.545 4 24c3 4.364 11.2 12 20 12s17-7.636 20-12c-3-5.455-11-14.182-20-9.818ZM24 21c-5.816 0-11.585 1.473-14.69 2.523c-.445.15-.43.762.018.904C12.318 25.375 17.92 27 24 27c5.98 0 11.5-1.66 14.525-2.57l.025-.008c.461-.138.487-.784.03-.936C35.45 22.439 29.747 21 24 21Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsMouthNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}