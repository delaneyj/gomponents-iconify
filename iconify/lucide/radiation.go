package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h0M7.5 4.2c-.3-.5-.9-.7-1.3-.4C3.9 5.5 2.3 8.1 2 11c-.1.5.4 1 1 1h5c0-1.5.8-2.8 2-3.4c-1.1-1.9-2-3.5-2.5-4.4zM21 12c.6 0 1-.4 1-1c-.3-2.9-1.8-5.5-4.1-7.1c-.4-.3-1.1-.2-1.3.3c-.6.9-1.5 2.5-2.6 4.3c1.2.7 2 2 2 3.5h5zM7.5 19.8c-.3.5-.1 1.1.4 1.3c2.6 1.2 5.6 1.2 8.2 0c.5-.2.7-.8.4-1.3c-.5-.9-1.4-2.5-2.5-4.3c-1.2.7-2.8.7-4 0c-1.1 1.8-2 3.4-2.5 4.3z"/>`),
		g.Group(children),
	)
}