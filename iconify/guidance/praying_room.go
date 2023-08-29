package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrayingRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m17 9l-4.25 4.5h-.25l-1.5-6h-.5a3 3 0 0 0-3 3V17c5.5 0 9 2.5 9 3.5H7m6-10.882L15 7.5m-15 16h24m-12.445-18s1.81-.557 2.135-1.776a1.768 1.768 0 0 0-1.242-2.163a1.75 1.75 0 0 0-2.146 1.25c-.324 1.219.963 2.61.963 2.61l.29.079Z"/>`),
		g.Group(children),
	)
}