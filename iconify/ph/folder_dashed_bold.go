package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDashedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 208a12 12 0 0 1-12 12H39.38A19.41 19.41 0 0 1 20 200.62V192a12 12 0 0 1 24 0v4h44a12 12 0 0 1 12 12Zm60-12h-32a12 12 0 0 0 0 24h32a12 12 0 0 0 0-24Zm64-56a12 12 0 0 0-12 12v44h-12a12 12 0 0 0 0 24h16.89A19.13 19.13 0 0 0 236 200.89V152a12 12 0 0 0-12-12Zm-8-72h-48a12 12 0 0 0 0 24h44v20a12 12 0 0 0 24 0V88a20 20 0 0 0-20-20ZM32 164a12 12 0 0 0 12-12v-32a12 12 0 0 0-24 0v32a12 12 0 0 0 12 12ZM20 80V52a20 20 0 0 1 20-20h52.41a20 20 0 0 1 14.94 6.71L137 72a12 12 0 0 1-9 20H32a12 12 0 0 1-12-12Zm24-12h57.28L90.61 56H44Z"/>`),
		g.Group(children),
	)
}