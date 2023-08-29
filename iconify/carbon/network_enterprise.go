package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkEnterprise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.798 10a10 10 0 0 0-19.62.124A7.496 7.496 0 0 0 7.5 25H8v-2h-.5a5.496 5.496 0 0 1-.377-10.98l.837-.057l.09-.833A7.993 7.993 0 0 1 23.736 10Z"/><path fill="currentColor" d="M28 12H18a2.002 2.002 0 0 0-2 2v4h-4a2.002 2.002 0 0 0-2 2v10h20V14a2.002 2.002 0 0 0-2-2ZM12 28v-8h4v8Zm16 0H18V14h10Z"/><path fill="currentColor" d="M20 16h2v4h-2zm4 0h2v4h-2zm-4 6h2v4h-2zm4 0h2v4h-2z"/>`),
		g.Group(children),
	)
}