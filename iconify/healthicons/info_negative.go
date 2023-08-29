package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsInfoNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm0-7a2 2 0 0 1-2-2V21a2 2 0 1 1 4 0v14a2 2 0 0 1-2 2Zm-2-24a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsInfoNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}