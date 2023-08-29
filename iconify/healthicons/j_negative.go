package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsJNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM32 12a2 2 0 1 0-4 0v18a4 4 0 1 1-8 0a2 2 0 1 0-4 0a8 8 0 1 0 16 0V12Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsJNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}