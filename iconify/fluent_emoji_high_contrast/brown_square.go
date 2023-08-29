package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrownSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 1a3 3 0 0 1 3 3v24a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3h24Zm1 3a1 1 0 0 0-1-1h-1.998L29 5.998V4Zm0 2.705L25.295 3h-4.293L29 10.998V6.705Zm0 5L20.295 3h-4.293L29 15.998v-4.293Zm0 5L15.295 3h-4.293L29 20.998v-4.293Zm0 5L10.295 3H6.002L29 25.998v-4.293Zm0 5L5.295 3H4a1 1 0 0 0-1 1v.998L27.002 29H28a1 1 0 0 0 1-1v-1.295ZM26.295 29L3 5.705v4.293L22.002 29h4.293Zm-5 0L3 10.705v4.293L17.002 29h4.293Zm-5 0L3 15.705v4.293L12.002 29h4.293Zm-5 0L3 20.705v4.293L7.002 29h4.293Zm-5 0L3 25.705V28a1 1 0 0 0 1 1h2.295Z"/>`),
		g.Group(children),
	)
}