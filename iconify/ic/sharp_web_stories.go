package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpWebStories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4h2v16h-2V4zM2 2v20h13V2H2zm19 16h1.5V6H21v12z"/>`),
		g.Group(children),
	)
}