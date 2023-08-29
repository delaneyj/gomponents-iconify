package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 88a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm-40-16a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm120 72a91.84 91.84 0 0 1-2.34 20.64l19.15 8.36a12 12 0 0 1-9.62 22l-18-7.85a92 92 0 0 1-162.46 0l-18 7.85a12 12 0 1 1-9.62-22l19.15-8.36A91.84 91.84 0 0 1 36 144v-4H16a12 12 0 0 1 0-24h20v-4a91.84 91.84 0 0 1 2.34-20.64L19.19 83a12 12 0 0 1 9.62-22l18 7.85a92 92 0 0 1 162.46 0l18-7.85a12 12 0 1 1 9.62 22l-19.15 8.36A91.84 91.84 0 0 1 220 112v4h20a12 12 0 0 1 0 24h-20ZM60 116h136v-4a68 68 0 0 0-136 0Zm56 94.92V140H60v4a68.1 68.1 0 0 0 56 66.92ZM196 144v-4h-56v70.92A68.1 68.1 0 0 0 196 144Z"/>`),
		g.Group(children),
	)
}