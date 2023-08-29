package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.001 17a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Zm0-2a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm-13-13h5v20h-5V2Zm2 2v16h1V4h-1Z"/>`),
		g.Group(children),
	)
}