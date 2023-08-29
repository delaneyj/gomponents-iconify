package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsKNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 12v11.172l12.586-12.586a2 2 0 1 1 2.828 2.828L22.828 24l10.586 10.586a2 2 0 1 1-2.828 2.828L20 26.828l-2 2V36a2 2 0 1 1-4 0V12a2 2 0 1 1 4 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsKNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}