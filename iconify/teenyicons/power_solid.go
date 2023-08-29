package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 0v9H7V0h1Zm4.387 1.792l.358.35A7.468 7.468 0 0 1 15 7.494C15 11.635 11.644 15 7.5 15C3.358 15 0 11.635 0 7.495a7.46 7.46 0 0 1 2.269-5.354l.357-.35l.7.716l-.359.35A6.463 6.463 0 0 0 1 7.494A6.506 6.506 0 0 0 7.5 14c3.59 0 6.5-2.917 6.5-6.505a6.464 6.464 0 0 0-1.955-4.639l-.357-.35l.7-.714Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}