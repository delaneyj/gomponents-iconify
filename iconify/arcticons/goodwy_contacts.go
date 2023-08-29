package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoodwyContacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.027" cy="24.013" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M27.373 24.937c5.604 1.056 7.926 3.89 7.926 3.89c.962 1.667 0 3.96 0 3.96c-3.25 4.046-8.181 4.471-11.297 4.471c-3.115 0-8.046-.425-11.296-4.471c0 0-.962-2.293 0-3.96c0 0 2.322-2.833 7.926-3.89"/><ellipse cx="24.002" cy="18.369" rx="6.64" ry="7.625"/></g>`),
		g.Group(children),
	)
}