package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BanknoteTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 11c0-2.828 0-4.243.879-5.121C3.757 5 5.172 5 8 5h5c2.828 0 4.243 0 5.121.879C19 6.757 19 8.172 19 11c0 2.828 0 4.243-.879 5.121C17.243 17 15.828 17 13 17H8c-2.828 0-4.243 0-5.121-.879C2 15.243 2 13.828 2 11Z"/><path d="M19 8.076c.975.096 1.631.313 2.121.803c.88.878.88 2.293.88 5.121c0 2.828 0 4.243-.88 5.121c-.878.88-2.293.88-5.12.88h-5c-2.83 0-4.244 0-5.122-.88c-.49-.49-.707-1.146-.803-2.121"/><path d="M13 11a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path stroke-linecap="round" d="M16 13V9M5 13V9"/></g>`),
		g.Group(children),
	)
}