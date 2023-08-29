package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderFollow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4.5h12V11h1.5V3h-15v12a2 2 0 0 0 2 2h7v-1.5h-7A.5.5 0 0 1 4 15zm10.5 2h-9V8h9zm-5 3h-4V11h4zM13 11h-1v1h1zm-1-1.5h-1.5v4h4v-4H13zM9.5 12h-4v1.5h4zm6.5 1.25h1.5v2.25h2.25V17H17.5v2.25H16V17h-2.25v-1.5H16z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}