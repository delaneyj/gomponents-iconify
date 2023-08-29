package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomPan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M27.01 12l-1.41 1.41L28.18 16l-2.59 2.59L27.01 20l4-4l-4-4z" fill="currentColor"/><path d="M6.41 13.42L5 12l-4 4l4 4l1.42-1.41L3.83 16l2.58-2.58z" fill="currentColor"/><path d="M16 28.17l-2.59-2.59L12 27l4 4l4-4l-1.41-1.41L16 28.17z" fill="currentColor"/><path d="M16 3.83l2.58 2.58L20 5l-4-4l-4 4l1.41 1.42L16 3.83z" fill="currentColor"/><path d="M22 16a6 6 0 1 0-2.53 4.89l3.82 3.82l1.42-1.42l-3.82-3.82A6 6 0 0 0 22 16zm-6 4a4 4 0 1 1 4-4a4 4 0 0 1-4 4z" fill="currentColor"/>`),
		g.Group(children),
	)
}