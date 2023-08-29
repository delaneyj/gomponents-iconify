package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Four(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M28.58 10.086A2 2 0 0 1 30 12v16h2a2 2 0 1 1 0 4h-2v4a2 2 0 0 1-4 0v-4H16a2 2 0 0 1-1.664-3.11l12-18a2 2 0 0 1 2.244-.804ZM26 28v-9.394L19.737 28H26Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}