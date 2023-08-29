package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 3.636c0-.876.242-1.524.642-1.95c.395-.421 1.001-.686 1.86-.686c.946 0 1.582.306 1.97.806c.331.427.528 1.037.528 1.827h1c0-.95-.237-1.794-.738-2.44C13.64.39 12.674 0 11.502 0c-1.073 0-1.967.338-2.59 1.002C8.294 1.662 8 2.582 8 3.636V6H1a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H9V3.636ZM1 7h9v6H1V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}