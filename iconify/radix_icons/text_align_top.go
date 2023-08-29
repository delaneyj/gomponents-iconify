package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.9 9.5a.4.4 0 0 1-.8 0V2.466L1.781 3.783a.4.4 0 0 1-.565-.566l2-2a.4.4 0 0 1 .565 0l2 2a.4.4 0 0 1-.565.566L3.899 2.466V9.5ZM8.5 2a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1h-6Zm0 3a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1h-6ZM8 8.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}