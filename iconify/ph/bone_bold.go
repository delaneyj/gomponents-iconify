package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoneBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M234.74 58.31a39.89 39.89 0 0 0-23.35-13.7a40 40 0 1 0-74.86 25.87l-66 66.05a40 40 0 1 0-25.87 74.86a40 40 0 1 0 74.86-25.86l66.05-66.06a40 40 0 0 0 49.22-61.16Zm-18.64 36.2a16 16 0 0 1-20.54 3a20 20 0 0 0-24.84 2.76l-70.4 70.4a20 20 0 0 0-2.75 24.84a16 16 0 1 1-29.41 6A12 12 0 0 0 56.3 187.7a11.61 11.61 0 0 0-1.85.14a16 16 0 1 1 6-29.4a20 20 0 0 0 24.83-2.76l70.4-70.4a20 20 0 0 0 2.76-24.84a16 16 0 1 1 29.4-6a12 12 0 0 0 13.71 13.71a16 16 0 0 1 14.55 26.36Z"/>`),
		g.Group(children),
	)
}