package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 4.5h17v5h.054c.94 0 1.877-.274 2.564-.917a10.929 10.929 0 0 0 2.061-2.624L23 4.5h.5v15H23l-.82-1.459a10.928 10.928 0 0 0-2.062-2.624c-.687-.643-1.624-.917-2.564-.917H17.5v5H.5v-15Z"/>`),
		g.Group(children),
	)
}