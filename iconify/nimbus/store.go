package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Store(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2.49 12h4.33V7.66H2.49zm1.25-3.09h1.83v1.83H3.74zm7.33-1.25a2.43 2.43 0 0 0-2.43 2.43v3.4H9.9v-3.4a1.18 1.18 0 1 1 2.35 0v3.4h1.25v-3.4a2.43 2.43 0 0 0-2.43-2.43zM2.49 5.07H13.5v1.3H2.49z"/><path fill="currentColor" d="M14.12 2.51H1.88A1.88 1.88 0 0 0 0 4.39v9.1h1.25v-9.1a.63.63 0 0 1 .63-.63h12.24a.63.63 0 0 1 .63.63v9.1H16v-9.1a1.88 1.88 0 0 0-1.88-1.88z"/>`),
		g.Group(children),
	)
}