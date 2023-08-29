package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBookSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V4H1v1h1v5H1v1h1v2.5A1.5 1.5 0 0 0 3.5 15h9a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 12.5 0h-9ZM6 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm-1 6a3 3 0 1 1 6 0v1H5v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}