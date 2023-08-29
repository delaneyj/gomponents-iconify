package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3a1 1 0 1 0-2 0v.75h-.557A4.214 4.214 0 0 0 6.237 7.7l-.221 3.534a7.377 7.377 0 0 1-1.308 3.754a1.61 1.61 0 0 0 .017 1.872c.081.112.242.112.34.014L16.574 5.365a.24.24 0 0 0 .007-.336a4.205 4.205 0 0 0-3.024-1.279H13V3Z"/><path fill="currentColor" fill-rule="evenodd" d="M17.786 8.074L21.03 4.83a.75.75 0 0 0-1.06-1.06l-16 16a.75.75 0 1 0 1.06 1.06l3.046-3.045l1.174.14V19a2.75 2.75 0 1 0 5.5 0v-1.075l3.407-.409a1.617 1.617 0 0 0 1.135-2.528a7.376 7.376 0 0 1-1.308-3.754l-.198-3.16ZM10.75 19a1.25 1.25 0 1 0 2.5 0v-.75h-2.5V19Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}