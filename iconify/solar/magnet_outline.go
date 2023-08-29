package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 2.75a9.25 9.25 0 0 0 0 18.5h3.25v-3.5H13a5.75 5.75 0 0 1 0-11.5h3.25v-3.5H13Zm4.75 0v3.5h1.75a.75.75 0 0 0 .75-.75v-2a.75.75 0 0 0-.75-.75h-1.75Zm0 15v3.5h1.75a.75.75 0 0 0 .75-.75v-2a.75.75 0 0 0-.75-.75h-1.75ZM2.25 12C2.25 6.063 7.063 1.25 13 1.25h6.5a2.25 2.25 0 0 1 2.25 2.25v2a2.25 2.25 0 0 1-2.25 2.25H13a4.25 4.25 0 0 0 0 8.5h6.5a2.25 2.25 0 0 1 2.25 2.25v2a2.25 2.25 0 0 1-2.25 2.25H13C7.063 22.75 2.25 17.937 2.25 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}