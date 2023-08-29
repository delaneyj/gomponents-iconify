package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M81 0h434c45 0 81 37 81 81v475c0 43-35 79-78 81v-54c13-2 24-14 24-27V81c0-14-13-26-27-26H81c-14 0-27 12-27 26c5-1 12-2 18-2c5 0 9 0 14 1l325 49c44 8 76 48 76 92v423c0 42-30 73-71 73c-4 0-10 0-15-1L77 667c-43-8-77-48-77-91V81C0 37 37 0 81 0zm16 255l285 42l7-49l-285-42zm60 43l-4 30l158 24l4-30z"/>`),
		g.Group(children),
	)
}