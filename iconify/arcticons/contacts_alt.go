package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5a9.237 9.237 0 1 1-9.22 9.254v-.017A9.237 9.237 0 0 1 24 5.5Zm0 21.897c10.32 0 18.457 6.873 18.457 10.296V42.5H5.543v-4.893c0-3.423 8.136-10.21 18.457-10.21Z"/>`),
		g.Group(children),
	)
}