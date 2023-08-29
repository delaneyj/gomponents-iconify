package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="80" height="80" x="336" y="336" fill="currentColor" rx="8" ry="8"/><rect width="64" height="64" x="272" y="272" fill="currentColor" rx="8" ry="8"/><rect width="64" height="64" x="416" y="416" fill="currentColor" rx="8" ry="8"/><rect width="48" height="48" x="432" y="272" fill="currentColor" rx="8" ry="8"/><rect width="48" height="48" x="272" y="432" fill="currentColor" rx="8" ry="8"/><rect width="80" height="80" x="336" y="96" fill="currentColor" rx="8" ry="8"/><rect width="176" height="176" x="288" y="48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="16" ry="16"/><rect width="80" height="80" x="96" y="96" fill="currentColor" rx="8" ry="8"/><rect width="176" height="176" x="48" y="48" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="16" ry="16"/><rect width="80" height="80" x="96" y="336" fill="currentColor" rx="8" ry="8"/><rect width="176" height="176" x="48" y="288" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" rx="16" ry="16"/>`),
		g.Group(children),
	)
}