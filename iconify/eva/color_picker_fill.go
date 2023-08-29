package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPickerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaColorPickerFill0"><g id="evaColorPickerFill1"><path id="evaColorPickerFill2" fill="currentColor" d="M19.4 7.34L16.66 4.6A1.92 1.92 0 0 0 14 4.53l-2 2l-1.29-1.24a1 1 0 0 0-1.42 1.42L10.53 8L5 13.53a2 2 0 0 0-.57 1.21L4 18.91a1 1 0 0 0 .29.8A1 1 0 0 0 5 20h.09l4.17-.38a2 2 0 0 0 1.21-.57l5.58-5.58l1.24 1.24a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-1.24-1.24l2-2a1.92 1.92 0 0 0-.07-2.71Zm-13 7.6L12 9.36l2.69 2.7l-2.79 2.79"/></g></g>`),
		g.Group(children),
	)
}