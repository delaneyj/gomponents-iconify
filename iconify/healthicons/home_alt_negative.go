package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHomeAltNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm20 33v8a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1V21.925l-3.294 3.283a1 1 0 0 1-1.412-1.416l17.056-17l.708-.706l.706.708l16.944 17a1 1 0 0 1-1.416 1.412L37 21.903V41a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-8a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHomeAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}