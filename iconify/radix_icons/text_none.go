package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.354 2.354a.5.5 0 0 0-.708-.708l-.717.718a.45.45 0 0 0-.429-.314h-8a.45.45 0 0 0-.45.45v2a.45.45 0 1 0 .9 0V2.95h3v4.393l-5.304 5.303a.5.5 0 0 0 .708.708L6.95 8.757v3.293H5.754a.45.45 0 1 0 0 .9h3.5a.45.45 0 0 0 0-.9H8.05V7.657l5.304-5.303ZM8.05 6.243l3-3V2.95h-3v3.293Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}