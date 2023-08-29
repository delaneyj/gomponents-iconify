package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Books(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.5 23A1.5 1.5 0 0 1 2 21.5V6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v1h4a2 2 0 0 1 2 2v2h5a2 2 0 0 1 2 2v15H16.5a.5.5 0 0 0 0 1h13.415a1.5 1.5 0 0 1-1.415 1h-12a1.5 1.5 0 0 1-1.5-1.5V26H9.5A1.5 1.5 0 0 1 8 24.5V23H3.5Zm13.415-1a1.5 1.5 0 0 1-1.415 1H22V9a1 1 0 0 0-1-1h-4v13H3.5a.5.5 0 0 0 0 1h13.415ZM9 23h.5c-.175 0-.344.03-.5.085V23Zm7-3V6a1 1 0 0 0-1-1H5v15h11Zm0 6v1.085c.156-.055.325-.085.5-.085H29V13a1 1 0 0 0-1-1h-5v12H9.5a.5.5 0 0 0 0 1h13.415a1.5 1.5 0 0 1-1.415 1H16Z"/>`),
		g.Group(children),
	)
}