package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Container(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM5 13h5V2H5v11Zm-1 0a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v11Zm9.5-11a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 5.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 7.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 9.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 11.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm11.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}