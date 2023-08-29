package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmotionSadLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10c0 .727-.078 1.435-.225 2.118l-1.782-1.783a8 8 0 1 0-4.374 6.801a3.998 3.998 0 0 0 1.555 1.423A9.955 9.955 0 0 1 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2Zm7 12.172l1.414 1.414a2 2 0 1 1-2.93.11l.102-.11L19 14.172ZM12 15c1.466 0 2.785.631 3.7 1.637l-.945.86C13.965 17.182 13.018 17 12 17c-1.018 0-1.965.183-2.755.496l-.945-.86A4.987 4.987 0 0 1 12 15Zm-3.5-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm7 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/>`),
		g.Group(children),
	)
}