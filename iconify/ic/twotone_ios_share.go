package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneIosShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8h-3v2h3v11H6V10h3V8H6c-1.11 0-2 .89-2 2v11a2 2 0 0 0 2 2h12c1.1 0 2-.9 2-2V10a2 2 0 0 0-2-2z"/><path fill="currentColor" d="M11 16h2V5h3l-4-4l-4 4h3z"/>`),
		g.Group(children),
	)
}