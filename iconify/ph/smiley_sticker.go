package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileySticker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174.92 156c-10.29 17.79-27.39 28-46.92 28s-36.63-10.2-46.93-28a8 8 0 1 1 13.86-8c7.46 12.91 19.2 20 33.07 20s25.61-7.1 33.08-20a8 8 0 1 1 13.84 8ZM232 128a104.35 104.35 0 0 1-4.56 30.56a8 8 0 0 1-2 3.31l-63.57 63.57a7.9 7.9 0 0 1-3.3 2A104 104 0 1 1 232 128Zm-16 0a87.89 87.89 0 1 0-64 84.69L212.69 152a88.05 88.05 0 0 0 3.31-24Zm-124-8a12 12 0 1 0-12-12a12 12 0 0 0 12 12Zm72-24a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}