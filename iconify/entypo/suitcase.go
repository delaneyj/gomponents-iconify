package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 4h-1v15h1c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zM0 6v11c0 1.1.899 2 2 2h1V4H2C.899 4 0 4.9 0 6zm13.5-4.094C12.819 1.59 11.611 1 9.981 1c-1.633 0-2.8.59-3.481.906V4H4v15h12V4h-2.5V1.906zM12 4H8V2.664c.534-.23 1.078-.465 1.981-.465c.902 0 1.486.234 2.019.465V4z"/>`),
		g.Group(children),
	)
}