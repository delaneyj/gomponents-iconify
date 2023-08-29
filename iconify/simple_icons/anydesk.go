package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anydesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.322 3.677L0 12l8.322 8.323L16.645 12zm7.371.01l-1.849 1.85l6.49 6.456l-6.49 6.49l1.85 1.817L24 11.993Z"/>`),
		g.Group(children),
	)
}