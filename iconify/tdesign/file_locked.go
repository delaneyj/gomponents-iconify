package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6.5v2H3V1Zm12 2.414V7h3.586L15 3.414ZM18 14.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25Zm3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75h-1.251V23h9v-6.5H21.25Zm-.751 2V21h-5v-2.5h5Z"/>`),
		g.Group(children),
	)
}