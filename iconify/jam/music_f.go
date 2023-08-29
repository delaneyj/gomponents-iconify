package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.954 12.87V1.564a1 1 0 0 1 1-1h.028a1 1 0 0 1 1 1V15.63c.197 1.969-1.42 3.99-3.874 4.693c-2.69.772-5.368-.333-5.98-2.468c-.612-2.135 1.073-4.491 3.764-5.263c1.47-.421 2.935-.283 4.062.277z"/>`),
		g.Group(children),
	)
}