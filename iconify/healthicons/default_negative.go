package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DefaultNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsDefaultNegative0)"><path d="M13 13a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2H13Zm0 5a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2H13Zm-1 6a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2h-5Zm22.981 2.293a1 1 0 0 0-1.414 0l-.791.79l1.414 1.415l.791-.79a1 1 0 0 0 0-1.415Zm-2.912 1.497L28 36.86v1.414h1.414l5.07-5.069l-1.415-1.414Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm8 9v26a3 3 0 0 0 3 3h11.832A9.966 9.966 0 0 1 22 34c0-5.523 4.477-10 10-10c.685 0 1.354.069 2 .2V9a3 3 0 0 0-3-3H11a3 3 0 0 0-3 3Zm32 25a8 8 0 1 1-16 0a8 8 0 0 1 16 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDefaultNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}