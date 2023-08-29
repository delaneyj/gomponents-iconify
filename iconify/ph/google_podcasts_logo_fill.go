package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePodcastsLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m243.32 116.68l-104-104a16 16 0 0 0-22.64 0l-104 104a16 16 0 0 0 0 22.64l104 104a16 16 0 0 0 22.64 0l104-104a16 16 0 0 0 0-22.64ZM56 136a8 8 0 0 1-16 0v-16a8 8 0 0 1 16 0Zm40 40a8 8 0 0 1-16 0v-16a8 8 0 0 1 16 0Zm0-48a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Zm40 88a8 8 0 0 1-16 0v-16a8 8 0 0 1 16 0Zm0-48a8 8 0 0 1-16 0V88a8 8 0 0 1 16 0Zm0-112a8 8 0 0 1-16 0V40a8 8 0 0 1 16 0Zm40 120a8 8 0 0 1-16 0v-48a8 8 0 0 1 16 0Zm0-80a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0Zm40 40a8 8 0 0 1-16 0v-16a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}