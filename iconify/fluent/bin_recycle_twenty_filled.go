package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BinRecycleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 3h9a1 1 0 0 1 1 1h-11a1 1 0 0 1 1-1Zm-2 1v.5c0 .022.001.044.004.065l1.347 11.664A2 2 0 0 0 6.837 18h6.327a2 2 0 0 0 1.986-1.77l1.346-11.666a.5.5 0 0 0 .004-.049V4a2 2 0 0 0-2-2h-9a2 2 0 0 0-2 2Zm7.799 3.75l.452.782a.5.5 0 1 1-.867.5l-.451-.782a.5.5 0 0 0-.866 0l-.452.782a.5.5 0 1 1-.866-.5l.452-.782c.577-1 2.02-1 2.598 0Zm.866 3.5l-.108-.186a.5.5 0 0 1 .867-.5l.107.186A1.5 1.5 0 0 1 11.732 13H11a.5.5 0 0 1 0-1h.732a.5.5 0 0 0 .433-.75ZM9 12a.5.5 0 0 1 0 1h-.732a1.5 1.5 0 0 1-1.3-2.25l.108-.186a.5.5 0 0 1 .866.5l-.107.186a.5.5 0 0 0 .433.75H9Z"/>`),
		g.Group(children),
	)
}