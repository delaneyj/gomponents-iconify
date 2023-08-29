package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessPgnMaster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.61 14.45v3.49m0-3.49v-1.73h.65h-1.29h.64m1.64 1.73h-3.27"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM24 5.5v37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.482 38.483a.3.3 0 0 0 .3-.3v-.02c-.05-1.87-2.5-2-1.25-2.16c.65-.1 1.68-.85.85-1.71s-1.22-.67-.87-1s2-1.95.92-2.29c-1.6-.51-1-.7-.92-.77a2.9 2.9 0 0 0 1.61-2.28a.27.27 0 0 0-.48-.2c-.4.46-4 1.73-9.42 1.7c-.76.11-.37.91-.28 1.8c.06.54.9.65 1.28.61c.2 0 1.21-.13 1.36-.11c.64.08.79.57 1 1a2 2 0 0 0 1.92.88c.13 0-.25.15-1.13.14c-1.52 0-3.45 1.73-4.06 4.19a.39.39 0 0 0 .3.462a.396.396 0 0 0 .05.008c1.54.09 7.27.03 8.82.05Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.452 30.773a2.12 2.12 0 0 0 .94-.15a3.6 3.6 0 0 0-.94.15ZM14.58 20.35h-4.27s-3.87-7 4.25-11.78h0c8.12 4.82 4.26 11.78 4.26 11.78ZM5.5 24h37"/>`),
		g.Group(children),
	)
}