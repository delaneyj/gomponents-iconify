package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeSanitizerNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeSanitizerNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm13.928 9.048a8.748 8.748 0 0 0-1.56 1.948l-1.736-.992a10.736 10.736 0 0 1 1.94-2.427C13.43 6.787 14.624 6 16 6h17v2h-8v6h4a2 2 0 0 1 2 2v3h1a5 5 0 0 1 5 5v16a2 2 0 0 1-2 2H13a2 2 0 0 1-2-2V24a5 5 0 0 1 5-5h1v-3a2 2 0 0 1 2-2h4V8h-7c-.624 0-1.346.38-2.072 1.048ZM18 29h5v-5h2v5h5v2h-5v5h-2v-5h-5v-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeSanitizerNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}