package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletTrain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M19.85 7h10.154v11H30v12H2v-3h3.018c-2.506 0-3.92-2.876-2.376-4.85L6 18v9h7l1.234-1.596C15.406 23.888 17.237 23 19.191 23h9.814v-5H6l1.737-2h.004l.006-.007L12.961 10h-.002a9.43 9.43 0 0 1 6.892-3Zm-6.116 3.621l-.18.211L9.065 16h12.72a5.89 5.89 0 0 0 4.736-2.4l1.243-1.68c.582-.79.02-1.92-.957-1.92H15.223c-.61 0-1.098.172-1.489.621Z"/><path d="M8.99 22h2c.55 0 1-.45 1-1s-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1Z"/></g>`),
		g.Group(children),
	)
}