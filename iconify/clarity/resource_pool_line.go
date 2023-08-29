package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResourcePoolLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2ZM4 18a14 14 0 0 1 27.95-1H17.49L8.3 28.07A14 14 0 0 1 4 18Zm14 14a13.91 13.91 0 0 1-8.16-2.65L18.43 19h13.52A14 14 0 0 1 18 32Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}