package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func INegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsINegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16 10a2 2 0 1 0 0 4h6v20h-6a2 2 0 1 0 0 4h16a2 2 0 1 0 0-4h-6V14h6a2 2 0 1 0 0-4H16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsINegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}