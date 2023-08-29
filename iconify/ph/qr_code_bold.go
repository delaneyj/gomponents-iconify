package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 36H56a20 20 0 0 0-20 20v44a20 20 0 0 0 20 20h44a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 60H60V60h36Zm4 40H56a20 20 0 0 0-20 20v44a20 20 0 0 0 20 20h44a20 20 0 0 0 20-20v-44a20 20 0 0 0-20-20Zm-4 60H60v-36h36ZM200 36h-44a20 20 0 0 0-20 20v44a20 20 0 0 0 20 20h44a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 60h-36V60h36Zm-60 76v-24a12 12 0 0 1 24 0v24a12 12 0 0 1-24 0Zm84-8a12 12 0 0 1-12 12h-12v32a12 12 0 0 1-12 12h-36a12 12 0 0 1 0-24h24v-48a12 12 0 0 1 24 0v4h12a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}