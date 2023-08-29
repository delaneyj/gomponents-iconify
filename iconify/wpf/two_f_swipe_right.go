package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoFSwipeRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M15.906 3.969a1 1 0 0 0-.125.031a1 1 0 0 0-.5 1.719L16.562 7h-5.718C10.399 5.276 8.864 4 7 4a4 4 0 1 0 0 8c1.863 0 3.4-1.276 3.844-3h5.719l-1.282 1.281a1.016 1.016 0 1 0 1.438 1.438l3-3a1 1 0 0 0 0-1.438l-3-3a1 1 0 0 0-.813-.312zm4 10a1 1 0 0 0-.125.031a1 1 0 0 0-.5 1.719L20.563 17h-5.72c-.444-1.724-1.98-3-3.843-3a4 4 0 1 0 0 8c1.863 0 3.4-1.276 3.844-3h5.719l-1.282 1.281a1.016 1.016 0 1 0 1.438 1.438l3-3a1 1 0 0 0 0-1.438l-3-3a1 1 0 0 0-.813-.312z"/>`),
		g.Group(children),
	)
}