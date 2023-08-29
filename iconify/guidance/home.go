package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M22.5 21V10.25l-.247-.113a20 20 0 0 1-8.942-8.104L13 1.5h-2l-.311.533a20 20 0 0 1-8.942 8.104l-.247.113V21M23 22.5H1M15 21v-4a3 3 0 1 0-6 0v4"/>`),
		g.Group(children),
	)
}