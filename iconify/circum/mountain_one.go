package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.857 19.525l-6.57-14.96a2.5 2.5 0 0 0-4.58-.01l-6.56 14.96a1 1 0 0 0 .07.96a.985.985 0 0 0 .84.46h15.89a1 1 0 0 0 .91-1.41Zm-10.23-14.56a1.5 1.5 0 0 1 2.75 0l2.43 5.53l-1.45 1.45a.5.5 0 0 1-.71 0l-2.04-2.03a1.5 1.5 0 0 0-1.06-.44h-1.9Zm-6.57 14.96l4.15-9.45h2.34a.491.491 0 0 1 .36.15l2.03 2.03A1.508 1.508 0 0 0 14 13.1a1.491 1.491 0 0 0 1.06-.44l1.18-1.17l3.71 8.45Z"/>`),
		g.Group(children),
	)
}