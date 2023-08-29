package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="M49.212 51.534L67 35.036L49.212 18.538l-3.789 4.076l10.396 9.641H5v5.562h50.819l-10.396 9.642z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M49.212 51.534L67 35.036L49.212 18.538l-3.789 4.076l10.396 9.641H5v5.562h50.819l-10.396 9.642z"/>`),
		g.Group(children),
	)
}