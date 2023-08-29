package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFaceWithClosedEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ffdd67"/><path fill="#664e27" d="M45.6 45.4c1.3-1.9-2.3-2.6-2.8-5.5s2.6-4.8.8-6.1c-2.2-1.6-6 .6-9-1.6c.4 2.1 2.6 4.1 5.9 3.5c0 0-2.1 1.3-1.5 4.8c.6 3.5 3 4 3 4c-3.3.6-4.7 3.2-4.3 5.3c2-2.8 6.4-2.1 7.9-4.4"/><path fill="#ff717f" d="M55 20.9c-4.2-1-8.2.4-9 3.2c-.8 2.8 1.9 6 6.1 7c4.2 1 8.2-.4 9-3.2c.8-2.8-2-5.9-6.1-7m-33.1 8.8c-2.4-2.9-8.3-3-13.1-.2C4 32.4 2 37.1 4.5 40c2.4 2.9 8.3 3 13.1.2c4.8-2.8 6.7-7.5 4.3-10.5" opacity=".8"/><path fill="#664e27" d="M35.2 17.7c5.5 7.1 13.9 5.6 16.7-2.9c.1-.4-.4-.5-1.2-.8c-3.6 4-10.4 4.9-14.7 2.6c-.6.4-1.1.7-.8 1.1M10.8 22c5.5 7.1 13.9 5.6 16.7-2.9c.1-.4-.4-.5-1.2-.8c-3.6 4-10.4 4.9-14.7 2.6c-.6.4-1.1.7-.8 1.1"/>`),
		g.Group(children),
	)
}