package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26 8l-4-4l-4 4l1.41 1.42L21 7.83V18h2V7.83l1.58 1.58L26 8zM12.59 22.58L11 24.17V14H9v10.17l-1.58-1.58L6 24l4 4l4-4l-1.41-1.42zM2 2h2v28H2zm26 0h2v28h-2zM15 2h2v4h-2zm0 8h2v4h-2zm0 8h2v4h-2zm0 8h2v4h-2z"/>`),
		g.Group(children),
	)
}