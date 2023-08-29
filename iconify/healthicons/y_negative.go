package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsYNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM35.536 13.28a2 2 0 1 0-3.072-2.56L24 20.876L15.537 10.72a2 2 0 1 0-3.073 2.56L22 24.724V36a2 2 0 0 0 4 0V24.724l9.536-11.444Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsYNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}