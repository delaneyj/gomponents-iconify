package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePhotosLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 124h-46.32A68 68 0 0 0 128 20a4 4 0 0 0-4 4v46.32A68 68 0 0 0 20 128a4 4 0 0 0 4 4h46.32A68 68 0 0 0 128 236a4 4 0 0 0 4-4v-46.32A68 68 0 0 0 236 128a4 4 0 0 0-4-4Zm-44-36a59.28 59.28 0 0 1-12 36h-44V28.13A60.08 60.08 0 0 1 188 88ZM88 68a59.28 59.28 0 0 1 36 12v44H28.13A60.08 60.08 0 0 1 88 68ZM68 168a59.28 59.28 0 0 1 12-36h44v95.87A60.08 60.08 0 0 1 68 168Zm100 20a59.28 59.28 0 0 1-36-12v-44h95.87A60.08 60.08 0 0 1 168 188Z"/>`),
		g.Group(children),
	)
}