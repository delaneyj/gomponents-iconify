package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardTextBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 156a12 12 0 0 1-12 12H96a12 12 0 0 1 0-24h64a12 12 0 0 1 12 12Zm-12-52H96a12 12 0 0 0 0 24h64a12 12 0 0 0 0-24Zm60-56v168a20 20 0 0 1-20 20H56a20 20 0 0 1-20-20V48a20 20 0 0 1 20-20h34.53a51.88 51.88 0 0 1 74.94 0H200a20 20 0 0 1 20 20Zm-92-12a28 28 0 0 0-27.71 24h55.42A28 28 0 0 0 128 36Zm68 16h-17.41A52.13 52.13 0 0 1 180 64v8a12 12 0 0 1-12 12H88a12 12 0 0 1-12-12v-8a52.13 52.13 0 0 1 1.41-12H60v160h136Z"/>`),
		g.Group(children),
	)
}