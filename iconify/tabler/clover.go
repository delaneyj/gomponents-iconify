package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10L8.603 6.56a2.104 2.104 0 0 1 0-2.95a2.04 2.04 0 0 1 2.912 0L12 4l.485-.39a2.04 2.04 0 0 1 2.912 0a2.104 2.104 0 0 1 0 2.95L12 10zm0 4l-3.397 3.44a2.104 2.104 0 0 0 0 2.95a2.04 2.04 0 0 0 2.912 0L12 20l.485.39a2.04 2.04 0 0 0 2.912 0a2.104 2.104 0 0 0 0-2.95L12 14zm2-2l3.44-3.397a2.104 2.104 0 0 1 2.95 0a2.04 2.04 0 0 1 0 2.912L20 12l.39.485a2.04 2.04 0 0 1 0 2.912a2.104 2.104 0 0 1-2.95 0L14 12zm-4 0L6.56 8.603a2.104 2.104 0 0 0-2.95 0a2.04 2.04 0 0 0 0 2.912L4 12l-.39.485a2.04 2.04 0 0 0 0 2.912a2.104 2.104 0 0 0 2.95 0L10 12z"/>`),
		g.Group(children),
	)
}