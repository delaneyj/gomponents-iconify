package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 20a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M16 24a10.655 10.655 0 0 1-9.97-7.758L5.97 16l.06-.242A10.655 10.655 0 0 1 16 8a10.655 10.655 0 0 1 9.97 7.758l.06.242l-.06.242A10.655 10.655 0 0 1 16 24Zm-7.965-8A8.598 8.598 0 0 0 16 22a8.598 8.598 0 0 0 7.965-6A8.598 8.598 0 0 0 16 10a8.598 8.598 0 0 0-7.965 6Z"/><path fill="currentColor" d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14Zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4Z"/>`),
		g.Group(children),
	)
}