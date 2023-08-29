package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsRNegative0)"><path d="M26 22h-8v-8h8a4 4 0 0 1 0 8Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM16 10a2 2 0 0 0-2 2v24a2 2 0 1 0 4 0V26h6.764l5.447 10.894a2 2 0 0 0 3.578-1.788l-4.835-9.67A8.003 8.003 0 0 0 26 10H16Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}