package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Support(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.222 19.778c4.296 4.296 11.26 4.296 15.556 0c4.296-4.296 4.296-11.26 0-15.556c-4.296-4.296-11.26-4.296-15.556 0c-4.296 4.296-4.296 11.26 0 15.556ZM6.343 21.9l4.243-4.242m-8.485 0l4.242-4.243m11.314-2.828l4.242-4.243m-8.485 0l4.243-4.242m-9.9 14.142a6 6 0 1 0 8.486-8.486a6 6 0 0 0-8.486 8.486Zm-5.656-9.9l4.242 4.243m0-8.485l4.243 4.242m2.828 11.314l4.243 4.242m0-8.485l4.242 4.243"/>`),
		g.Group(children),
	)
}