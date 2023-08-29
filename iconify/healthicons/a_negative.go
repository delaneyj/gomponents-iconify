package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ANegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsANegative0)"><path d="M24 17.2L28.5 28h-9L24 17.2Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM25.846 11.23a2 2 0 0 0-3.692 0L14.67 29.19a2.01 2.01 0 0 0-.035.084l-2.482 5.957a2 2 0 1 0 3.692 1.538L17.833 32h12.334l1.987 4.77a2 2 0 0 0 3.692-1.54l-2.482-5.956a1.937 1.937 0 0 0-.035-.085l-7.483-17.958Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsANegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}