package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512.342 1025q-118 0-222-50.5t-176-139.5l101-79q54 66 131.5 103.5t165.5 37.5q104 0 192.5-51.5t140-140t51.5-192.5t-51.5-192.5t-140-140t-192.5-51.5q-144 0-253 96l126 126q1 35-23 34h-330q-13 0-21.5-14t-9.5-29l-1-318q-1-24 34-23l134 134q70-63 158.5-98.5T512.342 1q104 0 199 41t163.5 109.5t109 163t40.5 198.5t-40.5 199t-109 163.5t-163.5 109t-199 40.5z"/>`),
		g.Group(children),
	)
}