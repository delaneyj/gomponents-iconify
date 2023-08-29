package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RdtResultOutStock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M27.707 6.293a1 1 0 0 1 0 1.414L25.414 10l2.293 2.293a1 1 0 0 1-1.414 1.414L24 11.414l-2.293 2.293a1 1 0 1 1-1.414-1.415L22.586 10l-2.293-2.293a1 1 0 1 1 1.415-1.414L24 8.586l2.293-2.293a1 1 0 0 1 1.414 0Zm-8.509 26.435A3.5 3.5 0 0 0 23 31.64v10.383L11 37.5v-7.504l8.198 2.732Z"/><path fill-rule="evenodd" d="m37 37.5l-12 4.523V31.64a3.5 3.5 0 0 0 3.802 1.088L37 29.996V37.5Zm-3.684-2.051a1 1 0 0 0-.632-1.898l-4.5 1.5a1 1 0 0 0 .632 1.898l4.5-1.5Zm-8.989-20.394a1 1 0 0 0-.654 0l-12.998 4.5l-.023.007a.996.996 0 0 0-.442.325l-3.99 4.988a1 1 0 0 0 .464 1.574l13.5 4.5a1 1 0 0 0 1.135-.376L24 26.743l2.68 3.83a1 1 0 0 0 1.136.376l13.5-4.5a1 1 0 0 0 .465-1.574l-3.99-4.988a.995.995 0 0 0-.466-.333l-12.998-4.499ZM24 23.942l9.943-3.442L24 17.058L14.057 20.5L24 23.942Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}