package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.955 3.106a.75.75 0 0 0-1.06-1.06l-.708.706A1.272 1.272 0 0 0 17.15 4.5a8.25 8.25 0 0 1-6.146 13.75h-.006c-1.968 0-3.935-.7-5.497-2.1a1.272 1.272 0 0 0-1.748.037l-.707.707a.75.75 0 1 0 1.06 1.061l.551-.55a9.709 9.709 0 0 0 5.594 2.316v1.53H8.5a.75.75 0 1 0 0 1.5h5a.75.75 0 0 0 0-1.5h-1.75v-1.53a9.712 9.712 0 0 0 6.144-2.827c3.63-3.629 3.8-9.407.51-13.238l.551-.55Z"/><path fill="currentColor" fill-rule="evenodd" d="M4.5 10a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0ZM13 6.25a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5ZM6.25 12.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}