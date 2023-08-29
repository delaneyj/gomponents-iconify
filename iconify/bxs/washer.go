package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Washer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22h16a1 1 0 0 0 1-1V5c0-1.654-1.346-3-3-3H6C4.346 2 3 3.346 3 5v16a1 1 0 0 0 1 1zM18 3.924a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm-3 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM12 7c3.309 0 6 2.691 6 6s-2.691 6-6 6s-6-2.691-6-6s2.691-6 6-6z"/><path fill="currentColor" d="M12.766 16.929c1.399-.261 2.571-1.315 3.023-2.665a3.853 3.853 0 0 0-.153-2.893a.482.482 0 0 0-.544-.266c-.604.149-1.019.448-1.5.801c-.786.577-1.765 1.294-3.592 1.294c-.813 0-1.45-.146-1.984-.354l-.013.009a4.006 4.006 0 0 0 4.763 4.074z"/>`),
		g.Group(children),
	)
}