package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 7.75h-1.4c.4-.48.65-1.08.65-1.75c0-1.52-1.23-2.75-2.75-2.75c-1.68 0-3.16.89-4 2.21a4.75 4.75 0 0 0-4-2.21C6.48 3.25 5.25 4.48 5.25 6c0 .67.25 1.27.65 1.75H4.5c-.69 0-1.25.56-1.25 1.25v2.5c0 .6.43 1.08 1 1.2v6.8c0 .69.56 1.25 1.25 1.25h13c.69 0 1.25-.56 1.25-1.25v-6.8c.57-.12 1-.6 1-1.2V9c0-.69-.56-1.25-1.25-1.25Zm-.25 3.5h-6.5v-2h6.5v2ZM16 4.75a1.25 1.25 0 0 1 0 2.5h-3.16c.34-1.43 1.63-2.5 3.16-2.5Zm-8 0c1.53 0 2.82 1.07 3.16 2.5H8a1.25 1.25 0 0 1 0-2.5Zm-3.25 4.5h6.5v2h-6.5v-2Zm1 3.5h5.5v6.5h-5.5v-6.5Zm12.5 6.5h-5.5v-6.5h5.5v6.5Z"/>`),
		g.Group(children),
	)
}