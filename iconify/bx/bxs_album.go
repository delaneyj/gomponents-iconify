package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsAlbum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm-4.607 8.055A4.956 4.956 0 0 0 7 12H5a6.978 6.978 0 0 1 2.051-4.95a6.978 6.978 0 0 1 2.225-1.5l.779 1.842c-.596.252-1.13.612-1.59 1.072s-.82.995-1.072 1.591zm4.6 3.945a2.007 2.007 0 1 1 0-4.014a2.007 2.007 0 0 1 0 4.014z" fill="currentColor"/>`),
		g.Group(children),
	)
}