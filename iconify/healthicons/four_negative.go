package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthicons4Negative0)"><path d="M26 18.606V28h-6.263L26 18.606Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM30 12a2 2 0 0 0-3.664-1.11l-12 18A2 2 0 0 0 16 32h10v4a2 2 0 0 0 4 0v-4h2a2 2 0 1 0 0-4h-2V12Z" clip-rule="evenodd"/></g><defs><clipPath id="healthicons4Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}