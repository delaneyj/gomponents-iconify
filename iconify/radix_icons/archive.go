package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.309 1a1 1 0 0 0-.894.553L1.053 4.276A.5.5 0 0 0 1 4.5V13a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1V4.5a.5.5 0 0 0-.053-.224l-1.362-2.723A1 1 0 0 0 11.691 1H3.309Zm0 1H7v2H2.309l1-2ZM8 4V2h3.691l1 2H8Zm-.5 1H13v8H2V5h5.5Zm-2 2a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}