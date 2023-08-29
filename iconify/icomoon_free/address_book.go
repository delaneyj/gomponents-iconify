package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0v16h12V0H3zm6 4.005a1.995 1.995 0 1 1 0 3.99a1.995 1.995 0 0 1 0-3.99zM12 12H6v-1a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v1zM1 1h1.5v3H1V1zm0 4h1.5v3H1V5zm0 4h1.5v3H1V9zm0 4h1.5v3H1v-3z"/>`),
		g.Group(children),
	)
}