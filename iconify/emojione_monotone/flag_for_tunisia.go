package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTunisia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M25.201 32c0-5.633 4.565-10.2 10.198-10.2a10.15 10.15 0 0 1 6.892 2.697c-2.318-3.175-6.059-5.246-10.291-5.246c-7.041 0-12.749 5.708-12.749 12.749S24.959 44.749 32 44.749c4.234 0 7.976-2.073 10.294-5.25a10.147 10.147 0 0 1-6.895 2.699c-5.632 0-10.198-4.565-10.198-10.198"/><path fill="currentColor" d="m36.302 29.221l-3.267-4.496v5.558L27.75 32l5.285 1.717v5.558l3.267-4.496l5.286 1.717L38.32 32l3.268-4.496z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 46.997c-9.388 0-16.999-7.609-16.999-16.997S22.612 15.001 32 15.001S48.997 22.612 48.997 32S41.388 48.997 32 48.997z"/>`),
		g.Group(children),
	)
}