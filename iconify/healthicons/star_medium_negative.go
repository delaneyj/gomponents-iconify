package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarMediumNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsStarMediumNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.892 11.06a.99.99 0 0 0-1.784 0l-3.556 7.278a.996.996 0 0 1-.749.55l-7.95 1.167c-.816.12-1.142 1.133-.551 1.714l5.752 5.665c.235.231.342.564.286.89l-1.358 7.999c-.139.82.714 1.447 1.444 1.06l7.11-3.777a.987.987 0 0 1 .927 0l7.11 3.776c.73.388 1.584-.238 1.445-1.06l-1.358-7.998a1.012 1.012 0 0 1 .286-.89l5.752-5.665c.591-.581.265-1.595-.551-1.714l-7.95-1.167a.996.996 0 0 1-.75-.55l-3.555-7.278Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsStarMediumNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}