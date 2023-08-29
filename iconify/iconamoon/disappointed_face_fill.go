package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisappointedFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm8.949-2.816a1 1 0 0 1-.633 1.265l-3 1a1 1 0 0 1-.632-1.898l3-1a1 1 0 0 1 1.265.633Zm3.367-.633a1 1 0 0 0-.632 1.898l3 1a1 1 0 0 0 .632-1.898l-3-1ZM9.5 14.67a1 1 0 1 0 1.002 1.73c.44-.254.95-.4 1.499-.4c.548 0 1.059.146 1.5.4a1 1 0 0 0 1-1.73A4.98 4.98 0 0 0 12 14c-.91 0-1.764.243-2.5.67Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}