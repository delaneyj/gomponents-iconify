package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23.1 6.849a1 1 0 0 1 1.8 0l16.4 33.714A1 1 0 0 1 40.403 42H7.599a1 1 0 0 1-.9-1.437L23.101 6.849ZM22 20a2 2 0 1 1 4 0v10a2 2 0 1 1-4 0V20Zm1.966 14a1.966 1.966 0 0 0 0 3.933h.068a1.966 1.966 0 1 0 0-3.933h-.067Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}