package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 10h-.1c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4V1H6v3.1C5.8 4 5.7 4 5.5 4C4.1 4 3 5.1 3 6.5c0 .6.2 1.1.6 1.6C3.4 8 3.2 8 3 8c-1.7 0-3 1.3-3 3s1.3 3 3 3h11c1.1 0 2-.9 2-2s-.9-2-2-2zm-6 1.4L5.1 8H7V2h2v6h1.9L8 11.4z"/>`),
		g.Group(children),
	)
}