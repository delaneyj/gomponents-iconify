package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodDropNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBloodDropNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 44c7.732 0 14-6.15 14-13.737C38 18.243 24 4 24 4S10 18.242 10 30.263C10 37.85 16.268 44 24 44Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBloodDropNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}