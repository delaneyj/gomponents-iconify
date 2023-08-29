package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillOneNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPill1Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm18-20c0 9.941-8.059 18-18 18S6 33.941 6 24S14.059 6 24 6s18 8.059 18 18Zm-7.793-9.707a1 1 0 0 1 0 1.414l-18.5 18.5a1 1 0 0 1-1.414-1.414l18.5-18.5a1 1 0 0 1 1.414 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPill1Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}