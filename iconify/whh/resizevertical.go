package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resizevertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M298 1015L9 736q-9-9-9-21.5T9 693l46-44q9-9 22.5-9t22.5 9l156 151V224L100 375q-9 9-22.5 9T55 375L9 331q-9-9-9-21.5T9 288L298 9q9-9 22.5-9T343 9l289 279q9 9 9 21.5t-9 21.5l-46 44q-9 9-22.5 9t-22.5-9L384 223v578l157-152q9-9 22.5-9t22.5 9l46 44q9 9 9 21.5t-9 21.5l-289 279q-9 9-22.5 9t-22.5-9z"/>`),
		g.Group(children),
	)
}