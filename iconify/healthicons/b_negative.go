package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsBNegative0)"><path d="M30 18a4 4 0 0 1-4 4h-8v-8h8a4 4 0 0 1 4 4Zm-4 8h-8v8h8a4 4 0 0 0 0-8Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM16 10a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h10a8 8 0 0 0 5.292-14A8 8 0 0 0 26 10H16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}