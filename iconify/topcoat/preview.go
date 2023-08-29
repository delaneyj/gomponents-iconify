package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M22.5 6.5c0-2.38-.5-3-3-3h-16c-2.5 0-3 .69-3 3v29c0 2.5.62 3 3 3h18.939L14 32.5H5.5v-24h12v20l5-4.561V6.5zm-9 30h-3v-2h3v2zm14.34-4.96a1.245 1.245 0 0 0 2.49 0c0-.69-.561-1.25-1.25-1.25c-.68 0-1.24.56-1.24 1.25zm2.61 7.021c4.05-.891 10.28-6.199 10.28-6.199S36.1 25.17 29.9 24.4c-.351-.041-1.57-.051-1.801-.03c-5.99.58-10.83 8.22-10.83 8.22s5.259 5.131 10.201 5.96c.58.12 2.38.12 2.98.011zm-7.4-7.032c0-3.1 2.671-5.609 5.95-5.609s5.95 2.51 5.95 5.609c0 3.111-2.671 5.621-5.95 5.621s-5.95-2.509-5.95-5.621z"/>`),
		g.Group(children),
	)
}