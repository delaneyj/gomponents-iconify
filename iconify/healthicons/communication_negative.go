package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunicationNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCommunicationNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM4 20.5C4 12.492 10.492 6 18.5 6h11C37.508 6 44 12.492 44 20.5S37.508 35 29.5 35H28v7S4 38.5 4 20.5Zm22.5.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm5.5 2.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM18.5 21a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCommunicationNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}