package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodAP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M20 10a1 1 0 0 1 .923.615l3.742 8.98l.017.042l1.241 2.978a1 1 0 0 1-1.846.77L23.083 21h-6.166l-.994 2.385a1 1 0 0 1-1.846-.77l1.241-2.978l.018-.042l3.74-8.98A1 1 0 0 1 20 10Zm2.25 9L20 13.6L17.75 19h4.5Z" clip-rule="evenodd"/><path d="M27 16a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2v-2a1 1 0 1 0-2 0v2h-2Z"/><path fill-rule="evenodd" d="M10 34.03c0 2.193 1.79 3.97 4 3.97h4v2h5v4h2v-4h5v-2h4c2.21 0 4-1.777 4-3.97V10.328c0-2.192-1.79-3.97-4-3.97h-6l-1.132-1.155a4.022 4.022 0 0 0-5.736 0L20 6.358h-6c-2.21 0-4 1.777-4 3.97V34.03ZM28 8.359a2 2 0 0 1-1.429-.6L25.44 6.602a2.021 2.021 0 0 0-2.878 0L21.43 7.758a2 2 0 0 1-1.429.6h-6c-1.12 0-2 .896-2 1.97V28.97a68.676 68.676 0 0 1 3.7.06l.53.02c1.284.048 2.481.093 3.652.069c2.627-.056 5.042-.466 7.61-1.981c3.118-1.84 5.758-1.288 7.583-.226c.338.197.647.41.925.626V10.327c0-1.073-.88-1.97-2-1.97h-6Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}