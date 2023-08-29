package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwordBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 28h-64a12 12 0 0 0-9.51 4.68l-61.23 79.6l-6.42-6.42a20 20 0 0 0-28.29 0l-16.69 16.69a20 20 0 0 0 0 28.29L45 166l-21.14 21.17a20 20 0 0 0 0 28.28l16.69 16.69a20 20 0 0 0 28.28 0L90 211l15.17 15.16a20 20 0 0 0 28.29 0l16.69-16.69a20 20 0 0 0 0-28.3l-6.42-6.41l79.6-61.23A12 12 0 0 0 228 104V40a12 12 0 0 0-12-12ZM54.69 212.34l-11-11L62 183l11 11Zm64.61-6L49.65 136.7l11.05-11l69.65 69.65ZM204 98.09l-77.39 59.53L121 152l47.51-47.5a12 12 0 0 0-17-17L104 135l-5.63-5.63L157.91 52H204Z"/>`),
		g.Group(children),
	)
}