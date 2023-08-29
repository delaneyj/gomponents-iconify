package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Importio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#001761" d="M167.785 56.146v156.411h-79.57V56.146C28.422 75.7-8.295 135.81 1.607 197.937C11.507 260.063 65.09 305.782 128 305.782c62.91 0 116.492-45.72 126.393-107.845c9.902-62.126-26.815-122.236-86.608-141.791Z"/><circle cx="128" cy="39.785" r="39.785" fill="#267CF9"/>`),
		g.Group(children),
	)
}