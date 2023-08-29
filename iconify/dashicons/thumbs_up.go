package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.72 2c.15-.02.26.02.41.07c.56.19.83.79.66 1.35c-.17.55-1 3.04-1 3.58c0 .53.75 1 1.35 1h3c.6 0 1 .4 1 1s-2 7-2 7c-.17.39-.55 1-1 1H6V8h2.14c.41-.41 3.3-4.71 3.58-5.27c.21-.41.6-.68 1-.73zM2 8h2v9H2V8z"/>`),
		g.Group(children),
	)
}