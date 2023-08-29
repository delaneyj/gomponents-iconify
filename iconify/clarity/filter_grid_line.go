package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterGridLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m15 25.86l2 1v-6.59a1 1 0 0 0-.29-.7L10.23 13h15.56l-6.47 6.57a1 1 0 0 0-.29.7L19 28l2 1v-8.32L27.58 14a1.46 1.46 0 0 0 .42-1v-1a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v1a1.46 1.46 0 0 0 .42 1L15 20.68Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}