package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Church(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="m49 22.396l-23-10.88V7h3V5h-3V1h-2v4h-3v2h3v4.516L7.942 19.121L1 22.396V24h7v25h12V34h10v15h12V24h7z"/>`),
		g.Group(children),
	)
}