package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 12V9H6a3 3 0 0 0-3 3v2a5 5 0 0 1 3-9h4V2l5 5l-5 5Z"/>`),
		g.Group(children),
	)
}