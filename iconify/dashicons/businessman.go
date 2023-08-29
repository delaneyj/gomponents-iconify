package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Businessman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 16.9v-2.5c0-.7-.1-1.4-.5-2.1c-.4-.7-.9-1.3-1.6-1.7c-.7-.5-2.2-.6-2.9-.6l-1.6 1.7l.6 1.3v3l-1 1.1L9 16v-3l.7-1.3L8 10c-.8 0-2.3.1-3 .6c-.7.4-1.1 1-1.5 1.7S3 13.6 3 14.4v2.5S5.6 18 10 18s7-1.1 7-1.1zM10 2.1c-1.9 0-3 1.8-2.7 3.8c.3 2 1.3 3.4 2.7 3.4s2.4-1.4 2.7-3.4c.3-2.1-.8-3.8-2.7-3.8z"/>`),
		g.Group(children),
	)
}