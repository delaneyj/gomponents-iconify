package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerPlayingCards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#DD2E44" d="M32 28H4V4a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v24z"/><path d="M8 36h20a4 4 0 0 0 4-4v-4c-4.117-2.744-21.139-8.233-28 0v4a4 4 0 0 0 4 4z"/><circle cx="15.276" cy="12.495" r="7.578" fill="#FFF"/>`),
		g.Group(children),
	)
}