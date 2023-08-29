package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#b0bdc6" d="m.6 48.7l18.5-19.3L4.9 8.9c-.7.9-1.2 2.1-1.3 3.5L0 45.1c-.1 1.4.1 2.6.6 3.6M62.7 5.5L35.9 30.4L57 58.9c1.2-1 2-2.5 2.2-4.3L64 10c.2-1.8-.4-3.5-1.3-4.5"/><path fill="#cad5dd" d="m25.7 39.4l-6.6-10L.6 48.7c.4.9 1.1 1.5 1.9 1.7L53.3 60c1.4.3 2.7-.2 3.7-1.1L35.9 30.4l-10.2 9"/><path fill="#dfe9ef" d="M59.3 4L7.1 7.5c-.8.1-1.6.6-2.2 1.4l14.3 20.5l6.6 10l10.2-9L62.7 5.5c-.9-1-2.1-1.6-3.4-1.5"/><path fill="#428bc1" d="m35.2 23.6l.6-5.1l-17.5-.1l-2.9 23.9l17 4l.7-5.1l-12.8-2.5l.7-5.5L33.8 35l.6-5.1l-12.9-1.3l.7-5.6z"/>`),
		g.Group(children),
	)
}