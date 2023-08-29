package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 14a1.99 1.99 0 0 0-1.981 2c0 1.104.887 2 1.981 2a1.99 1.99 0 0 0 1.98-2c0-1.105-.886-2-1.98-2zm-4.2-2.242l1.4 1.414a3.933 3.933 0 0 1 5.601 0l1.399-1.414a5.898 5.898 0 0 0-8.4 0zM3 8.928l1.4 1.414a7.864 7.864 0 0 1 11.199 0L17 8.928a9.831 9.831 0 0 0-14 0zM.199 6.1l1.4 1.414a11.797 11.797 0 0 1 16.801 0L19.8 6.1a13.763 13.763 0 0 0-19.601 0z"/>`),
		g.Group(children),
	)
}