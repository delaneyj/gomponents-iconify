package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 18.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm0-2a.501.501 0 1 1-.002 1.002A.501.501 0 0 1 17 16.5zm4.987-3.064a.488.488 0 0 0-.02-.095a.478.478 0 0 0-.043-.09a.481.481 0 0 0-.05-.073a.488.488 0 0 0-.08-.072c-.016-.012-.027-.029-.044-.039l-15.58-9a.5.5 0 1 0-.5.866L19.637 13H3.5a.5.5 0 0 0-.5.5v5A2.502 2.502 0 0 0 5.5 21h14a2.502 2.502 0 0 0 2.5-2.5v-5c0-.023-.01-.043-.013-.064zM21 18.5a1.5 1.5 0 0 1-1.5 1.5h-14A1.5 1.5 0 0 1 4 18.5V14h17v4.5z"/>`),
		g.Group(children),
	)
}