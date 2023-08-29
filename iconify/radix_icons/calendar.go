package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 1a.5.5 0 0 1 .5.5V2h5v-.5a.5.5 0 0 1 1 0V2h1.5A1.5 1.5 0 0 1 14 3.5v9a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-9A1.5 1.5 0 0 1 2.5 2H4v-.5a.5.5 0 0 1 .5-.5ZM10 3v.5a.5.5 0 0 0 1 0V3h1.5a.5.5 0 0 1 .5.5V5H2V3.5a.5.5 0 0 1 .5-.5H4v.5a.5.5 0 0 0 1 0V3h5ZM2 6v6.5a.5.5 0 0 0 .5.5h10a.5.5 0 0 0 .5-.5V6H2Zm5 1.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM9.5 7a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm.5 1.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM9 9.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM7.5 9a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM5 9.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0ZM3.5 9a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM3 11.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}