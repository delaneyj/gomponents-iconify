package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17 9.69v5.348l6.307-6.069c.635-.611 1.693-.16 1.693.72V22.31c0 .882-1.058 1.332-1.693.72L17 16.963v5.348c0 .882-1.058 1.332-1.693.72L9 16.963V22.5a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v5.538l6.307-6.069c.635-.611 1.693-.16 1.693.72Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}