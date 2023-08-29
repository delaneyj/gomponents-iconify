package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodOPNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBloodOPNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM28 8.358a2 2 0 0 1-1.429-.6L25.44 6.602a2.021 2.021 0 0 0-2.878 0L21.43 7.758a2 2 0 0 1-1.429.6h-6c-1.12 0-2 .896-2 1.97V28.97a68.676 68.676 0 0 1 3.7.06l.53.02c1.284.048 2.481.093 3.652.069c2.627-.056 5.042-.466 7.61-1.981c3.118-1.84 5.758-1.288 7.583-.226c.338.197.647.41.925.626V10.327c0-1.073-.88-1.97-2-1.97h-6ZM14 38h4v2h5v4h2v-4h5v-2h4c2.21 0 4-1.777 4-3.97V10.328c0-2.192-1.79-3.97-4-3.97h-6l-1.132-1.155a4.022 4.022 0 0 0-5.736 0L20 6.358h-6c-2.21 0-4 1.777-4 3.97V34.03c0 2.193 1.79 3.97 4 3.97Zm6-26a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm-4 6a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm11 0a1 1 0 0 1 1-1h1.5v-1.5a1 1 0 1 1 2 0V17H33a1 1 0 1 1 0 2h-1.5v1.5a1 1 0 1 1-2 0V19H28a1 1 0 0 1-1-1Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBloodOPNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}