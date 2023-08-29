package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quran(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 3C6.355 3 5 4.355 5 6v20c0 1.645 1.355 3 3 3h19V3zm0 2h17v18H8a2.95 2.95 0 0 0-1 .188V6c0-.566.434-1 1-1zm0 1v16h16V6zm2 2h12v12H10zm6 1c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1zm0 2h-3v6h6v-6zm3 3c0 .55.45 1 1 1s1-.45 1-1s-.45-1-1-1s-1 .45-1 1zm-3 3c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1zm-3-3c0-.55-.45-1-1-1s-1 .45-1 1s.45 1 1 1s1-.45 1-1zm2-1h2v2h-2zM8 25h17v2H8c-.566 0-1-.434-1-1c0-.566.434-1 1-1z"/>`),
		g.Group(children),
	)
}