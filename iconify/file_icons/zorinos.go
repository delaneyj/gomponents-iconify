package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zorinos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128.318 28.904L92.55 91.602h327.106l-35.355-62.698H128.318zM48.287 169.397L0 254.657l44.985 79.927h17.368l245.494-165.187H48.287zm399.02 0L201.78 334.584h265.372L512 255.344l-48.356-85.947h-16.336zM88.94 412.38l38.794 68.716h255.983l39.208-68.716H88.939z"/>`),
		g.Group(children),
	)
}