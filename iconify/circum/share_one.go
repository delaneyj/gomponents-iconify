package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.223 11.075a.5.5 0 0 0 .7.71l7-7v3.58a.508.508 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-4.79a.5.5 0 0 0-.5-.5h-4.79a.5.5 0 0 0 0 1h3.58Z"/><path fill="currentColor" d="M17.876 20.926H6.124a3.053 3.053 0 0 1-3.05-3.05V6.124a3.053 3.053 0 0 1 3.05-3.05h6.028a.5.5 0 0 1 0 1H6.124a2.053 2.053 0 0 0-2.05 2.05v11.752a2.053 2.053 0 0 0 2.05 2.05h11.752a2.053 2.053 0 0 0 2.05-2.05v-6.027a.5.5 0 0 1 1 0v6.027a3.053 3.053 0 0 1-3.05 3.05Z"/>`),
		g.Group(children),
	)
}