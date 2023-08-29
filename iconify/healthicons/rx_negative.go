package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RxNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsRxNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Zm-7 9a1 1 0 0 0-1 1v16h2v-7h2.586l6 6l-4.293 4.293l1.414 1.414L28 30.414l4.293 4.293l1.414-1.414L29.414 29l4.293-4.293l-1.414-1.414L28 27.586l-4.602-4.602A5 5 0 0 0 23 13h-6Zm1 8h5a3 3 0 1 0 0-6h-5v6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRxNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}