package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpIosShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8h-5v2h3v11H6V10h3V8H4v15h16z"/><path fill="currentColor" d="M11 16h2V5h3l-4-4l-4 4h3z"/>`),
		g.Group(children),
	)
}