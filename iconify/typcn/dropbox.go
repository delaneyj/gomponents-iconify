package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 12.9l5.3 3.5l3.7-3.1L6.7 10zm5.3-9.3L3 7.1L6.7 10L12 6.7zM21 7.1l-5.3-3.5L12 6.7l5.3 3.3zm-9 6.2l3.7 3.1l5.3-3.5l-3.7-2.9zm0 1.2l-3.7 3.1l-1.6-1.1v1.2l5.3 3.2l5.3-3.2v-1.2l-1.6 1.1z"/>`),
		g.Group(children),
	)
}