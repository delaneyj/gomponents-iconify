package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorefrontBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 96a12 12 0 0 0-.44-3.3L221.2 42.51A20.08 20.08 0 0 0 202 28H54a20.08 20.08 0 0 0-19.2 14.51L20.46 92.7A12 12 0 0 0 20 96v16a43.94 43.94 0 0 0 16 33.92V208a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20v-62.08A43.94 43.94 0 0 0 236 112Zm-24 16a20 20 0 0 1-40 0v-4h40Zm-168 0v-4h40v4a20 20 0 0 1-40 0Zm64-4h40v4a20 20 0 0 1-40 0ZM57.05 52H199l9.14 32H47.91ZM196 204H60v-48.19c1.32.12 2.65.19 4 .19a43.86 43.86 0 0 0 32-13.85a43.89 43.89 0 0 0 64 0A43.86 43.86 0 0 0 192 156c1.35 0 2.68-.07 4-.19Z"/>`),
		g.Group(children),
	)
}