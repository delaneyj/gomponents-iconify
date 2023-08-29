package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.35 4.5a2 2 0 0 0-1.95 2v35.1a2 2 0 0 0 1.95 2h27.3a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95ZM24 13.27a5.37 5.37 0 1 1-5.36 5.37A5.37 5.37 0 0 1 24 13.27ZM24 26c6 0 10.73 1.67 10.73 3.66v5.12H13.27v-5.17C13.27 27.62 18 26 24 26Z"/>`),
		g.Group(children),
	)
}