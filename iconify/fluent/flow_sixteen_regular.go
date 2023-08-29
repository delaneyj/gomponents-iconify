package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-2.45 2a2.5 2.5 0 1 0 0-1H9.5a2 2 0 0 0-2 2v2a1 1 0 0 1-1 1h-.55a2.5 2.5 0 1 0 0 1h.55a2 2 0 0 0 2-2V7a1 1 0 0 1 1-1h.55ZM5 10.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}