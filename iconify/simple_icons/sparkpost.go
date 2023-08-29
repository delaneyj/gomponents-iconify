package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkpost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.2 9c-1.351.9-1.8 2.7-1.65 3.9c-2.25-2.25 3.45-8.55-3-12.9C15.15 5.4 6 9.75 6 17.4c0 3 1.95 5.701 6 6.6c4.05-.898 6-3.6 6-6.6c0-4.5-2.7-6-1.8-8.4zM12 20.852c-1.8 0-3.45-1.5-3.45-3.451c0-1.801 1.5-3.45 3.45-3.45c1.8 0 3.45 1.5 3.45 3.45c-.15 1.951-1.65 3.451-3.45 3.451z"/>`),
		g.Group(children),
	)
}