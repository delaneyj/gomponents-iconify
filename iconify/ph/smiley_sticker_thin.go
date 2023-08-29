package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyStickerThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 29.39 195.61a3.94 3.94 0 0 0 1.66-1L222.62 159a4 4 0 0 0 1-1.65A100.07 100.07 0 0 0 128 28Zm88.25 126.1l-62.15 62.15a91.88 91.88 0 1 1 62.15-62.15ZM84 108a8 8 0 1 1 8 8a8 8 0 0 1-8-8Zm88 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-.54 46c-9.55 16.52-25.39 26-43.46 26s-33.91-9.48-43.46-26a4 4 0 0 1 6.92-4c8.21 14.19 21.19 22 36.54 22s28.33-7.81 36.54-22a4 4 0 1 1 6.92 4Z"/>`),
		g.Group(children),
	)
}