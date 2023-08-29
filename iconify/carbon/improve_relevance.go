package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImproveRelevance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 26.59L19.41 24L18 25.41l4 4l8-8L28.59 20L22 26.59z"/><circle cx="16" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M16 22a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Zm0-10a4 4 0 1 0 4 4a4.005 4.005 0 0 0-4-4Z"/><path fill="currentColor" d="M28 16a12 12 0 1 0-12 12v-2a10 10 0 1 1 10-10Z"/>`),
		g.Group(children),
	)
}