package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.25 6h23.5A6.25 6.25 0 0 1 42 12.25v23.5A6.25 6.25 0 0 1 35.75 42h-23.5A6.25 6.25 0 0 1 6 35.75v-23.5A6.25 6.25 0 0 1 12.25 6ZM8.5 16.5v19.25a3.75 3.75 0 0 0 3.75 3.75h23.5a3.75 3.75 0 0 0 3.75-3.75V16.5h-31Zm31-2.5v-1.75a3.75 3.75 0 0 0-3.75-3.75h-23.5a3.75 3.75 0 0 0-3.75 3.75V14h31Z"/>`),
		g.Group(children),
	)
}