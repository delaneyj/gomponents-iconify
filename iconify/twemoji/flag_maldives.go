package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMaldives(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#D21034" d="M32 5H4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4z"/><path fill="#007E3A" d="M6 11h24v14H6z"/><path fill="#FFF" d="M19.64 22.344c.279 0 .551-.027.814-.077a4.344 4.344 0 0 1 0-8.532a4.344 4.344 0 1 0-.814 8.609z"/>`),
		g.Group(children),
	)
}