package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamMagnetMountSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 16.3q-.825 0-1.413-.588T21 14.3v-4q0-.85.588-1.425T23 8.3h1v8h-1Zm-10-5.025q0 1-.375 1.913t-1.1 1.637L7.25 19.1L.175 12.05l4.275-4.3q.725-.725 1.638-1.1T8 6.275q2.1 0 3.55 1.45t1.45 3.55Zm3.95 4.075L15.9 14.3l1.2-1.2H14v-1.5h3.2l-1.25-1.25L17 9.3l3 3l-3.05 3.05Z"/>`),
		g.Group(children),
	)
}