package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.74 2.26a4.321 4.321 0 0 1 0 6.1l-1.53 1.52c-.32.33-.69.58-1.08.77L13 10l1.69-1.64l.76-.77l.76-.76c.84-.84.84-2.2 0-3.04a2.13 2.13 0 0 0-3.04 0l-.77.76l-.76.76L10 7l-.65-2.14c.19-.38.44-.75.77-1.07l1.52-1.53a4.321 4.321 0 0 1 6.1 0zM2 4l8 6l-6-8zm4-2l4 8l-2-8H6zM2 6l8 4l-8-2V6zm7.36 7.69L10 13l.74 2.35l-1.38 1.39a4.321 4.321 0 0 1-6.1 0a4.321 4.321 0 0 1 0-6.1l1.39-1.38L7 10l-.69.64l-1.52 1.53c-.85.84-.85 2.2 0 3.04c.84.85 2.2.85 3.04 0zM18 16l-8-6l6 8zm-4 2l-4-8l2 8h2zm4-4l-8-4l8 2v2z"/>`),
		g.Group(children),
	)
}