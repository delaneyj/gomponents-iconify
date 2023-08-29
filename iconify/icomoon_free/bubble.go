package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 1c4.418 0 8 2.91 8 6.5S12.418 14 8 14c-.424 0-.841-.027-1.247-.079c-1.718 1.718-3.77 2.027-5.753 2.072v-.421c1.071-.525 2-1.48 2-2.572a3.01 3.01 0 0 0-.034-.448C1.157 11.36 0 9.54 0 7.5C0 3.91 3.582 1 8 1z"/>`),
		g.Group(children),
	)
}