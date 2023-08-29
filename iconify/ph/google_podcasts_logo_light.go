package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePodcastsLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 16v32a6 6 0 0 1-12 0V16a6 6 0 0 1 12 0Zm42 42a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0V64a6 6 0 0 0-6-6Zm-48 144a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0v-32a6 6 0 0 0-6-6Zm0-120a6 6 0 0 0-6 6v80a6 6 0 0 0 12 0V88a6 6 0 0 0-6-6ZM80 58a6 6 0 0 0-6 6v56a6 6 0 0 0 12 0V64a6 6 0 0 0-6-6Zm96 72a6 6 0 0 0-6 6v56a6 6 0 0 0 12 0v-56a6 6 0 0 0-6-6ZM32 106a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0v-32a6 6 0 0 0-6-6Zm48 48a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0v-32a6 6 0 0 0-6-6Zm144-48a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0v-32a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}