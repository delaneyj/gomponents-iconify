package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostsCarouselVerticalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5.5h8a3 3 0 0 0 3-3a.5.5 0 0 0-.5-.5h-13a.5.5 0 0 0-.5.5a3 3 0 0 0 3 3Zm-3 6c0-1.886 0-2.828.586-3.414C6.172 7.5 7.114 7.5 9 7.5h6c1.886 0 2.828 0 3.414.586C19 8.672 19 9.614 19 11.5v1c0 1.886 0 2.828-.586 3.414c-.586.586-1.528.586-3.414.586H9c-1.886 0-2.828 0-3.414-.586C5 15.328 5 14.386 5 12.5v-1Zm11 7H8a3 3 0 0 0-3 3a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}