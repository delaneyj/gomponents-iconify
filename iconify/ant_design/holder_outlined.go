package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HolderOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M300 276.497a56 56 0 1 0 56-96.994a56 56 0 0 0-56 96.994Zm0 284a56 56 0 1 0 56-96.994a56 56 0 0 0-56 96.994ZM640 228a56 56 0 1 0 112 0a56 56 0 0 0-112 0Zm0 284a56 56 0 1 0 112 0a56 56 0 0 0-112 0ZM300 844.497a56 56 0 1 0 56-96.994a56 56 0 0 0-56 96.994ZM640 796a56 56 0 1 0 112 0a56 56 0 0 0-112 0Z"/>`),
		g.Group(children),
	)
}