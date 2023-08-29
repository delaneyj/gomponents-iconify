package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollHorizontalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 8a2 2 0 1 1 0 4H4a2 2 0 1 1 0-4h12Zm1 2a1 1 0 0 0-1-1H4a1 1 0 0 0 0 2h12a1 1 0 0 0 1-1ZM8 2a2 2 0 1 1 0 4H4a2 2 0 1 1 0-4h4Zm1 2a1 1 0 0 0-1-1H4a1 1 0 0 0 0 2h4a1 1 0 0 0 1-1Zm5 12a2 2 0 0 0-2-2H4a2 2 0 1 0 0 4h8a2 2 0 0 0 2-2Zm-2-1a1 1 0 1 1 0 2H4a1 1 0 1 1 0-2h8Z"/>`),
		g.Group(children),
	)
}