package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServersSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2.5A1.5 1.5 0 0 1 1.5 1h12A1.5 1.5 0 0 1 15 2.5v2c0 .384-.144.735-.382 1c.238.265.382.616.382 1v2c0 .384-.144.735-.382 1c.238.265.382.616.382 1v2a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-2c0-.384.144-.735.382-1A1.494 1.494 0 0 1 0 8.5v-2c0-.384.144-.735.382-1A1.494 1.494 0 0 1 0 4.5v-2ZM2 4h3V3H2v1Zm3 4H2V7h3v1Zm-3 4h3v-1H2v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}