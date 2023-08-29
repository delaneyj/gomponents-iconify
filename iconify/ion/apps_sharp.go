package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="112" height="112" x="48" y="48" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="200" y="48" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="352" y="48" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="48" y="200" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="200" y="200" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="352" y="200" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="48" y="352" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="200" y="352" fill="currentColor" rx="8" ry="8"/><rect width="112" height="112" x="352" y="352" fill="currentColor" rx="8" ry="8"/>`),
		g.Group(children),
	)
}