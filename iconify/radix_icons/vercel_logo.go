package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VercelLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.5 1l-.577 1.003L1.175 12L.6 13h13.8l-.575-1l-5.748-9.997L7.5 1Zm0 2.006L2.329 12H12.67L7.5 3.006Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}