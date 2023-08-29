package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YellowSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 1a3 3 0 0 0-3 3v24a3 3 0 0 0 3 3h24a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H4ZM3 4a1 1 0 0 1 1-1h1.998L3 5.998V4Zm0 2.705L6.705 3h4.293L3 10.998V6.705Zm0 5L11.705 3h4.293L3 15.998v-4.293Zm0 5L16.705 3h4.293L3 20.998v-4.293Zm0 5L21.705 3h4.293L3 25.998v-4.293Zm0 5L26.705 3H28a1 1 0 0 1 1 1v.998L4.998 29H4a1 1 0 0 1-1-1v-1.295ZM5.705 29L29 5.705v4.293L9.998 29H5.705Zm5 0L29 10.705v4.293L14.998 29h-4.293Zm5 0L29 15.705v4.293L19.998 29h-4.293Zm5 0L29 20.705v4.293L24.998 29h-4.293Zm5 0L29 25.705V28a1 1 0 0 1-1 1h-2.295Z"/>`),
		g.Group(children),
	)
}