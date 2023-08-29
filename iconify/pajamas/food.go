package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Food(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.75 1A5.75 5.75 0 0 0 1 6.75v.518a2 2 0 0 0 0 3.464v1.518A2.75 2.75 0 0 0 3.75 15h8.5A2.75 2.75 0 0 0 15 12.25v-1.518a2 2 0 0 0 0-3.464V6.75A5.75 5.75 0 0 0 9.25 1h-2.5ZM14 8.5H2a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1ZM13.5 7v-.25A4.25 4.25 0 0 0 9.25 2.5h-2.5A4.25 4.25 0 0 0 2.5 6.75V7h11ZM11 11h2.5v1.25c0 .69-.56 1.25-1.25 1.25h-8.5c-.69 0-1.25-.56-1.25-1.25V11H9l1 1l1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}