package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterGridSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M8 11v1.12a.5.5 0 0 0 .15.35l7.28 7.36a.5.5 0 0 1 .15.35v6.89a.5.5 0 0 0 .28.45l3.95 1.41a.5.5 0 0 0 .72-.45v-8.39a.54.54 0 0 1 .18-.35l7.12-7.25a.5.5 0 0 0 .15-.35V11Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}