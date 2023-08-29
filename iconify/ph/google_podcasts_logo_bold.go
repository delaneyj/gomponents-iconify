package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePodcastsLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 16v32a12 12 0 0 1-24 0V16a12 12 0 0 1 24 0Zm36 36a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0V64a12 12 0 0 0-12-12Zm-48 144a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0v-32a12 12 0 0 0-12-12Zm0-120a12 12 0 0 0-12 12v80a12 12 0 0 0 24 0V88a12 12 0 0 0-12-12ZM80 52a12 12 0 0 0-12 12v56a12 12 0 0 0 24 0V64a12 12 0 0 0-12-12Zm96 72a12 12 0 0 0-12 12v56a12 12 0 0 0 24 0v-56a12 12 0 0 0-12-12ZM32 100a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0v-32a12 12 0 0 0-12-12Zm48 48a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0v-32a12 12 0 0 0-12-12Zm144-48a12 12 0 0 0-12 12v32a12 12 0 0 0 24 0v-32a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}