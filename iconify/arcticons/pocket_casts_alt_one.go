package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PocketCastsAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.88 24H32a.4.4 0 0 1-.4-.38a7.63 7.63 0 1 0-8 8a.4.4 0 0 1 .4.38v3.85a.41.41 0 0 1-.41.41h0A12.29 12.29 0 1 1 36.3 23.57a.41.41 0 0 1-.39.43Zm8.8 0h-4.24a.8.8 0 0 1-.8-.77a15.64 15.64 0 1 0-16.4 16.41a.8.8 0 0 1 .77.8v4.27a.82.82 0 0 1-.82.8h0a21.51 21.51 0 1 1 22.32-22.34a.81.81 0 0 1-.79.83h0Z"/>`),
		g.Group(children),
	)
}