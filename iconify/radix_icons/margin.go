package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Margin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm3 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM8 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm3.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM1.5 14a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5-3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM1.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM2 4.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM13.5 11a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm.5-3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM13.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM5 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm3.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm2.5.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1ZM4 5a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5Zm1 0h5v5H5V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}