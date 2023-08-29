package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="M19.502 22.824L36 5.036l16.498 17.788l-4.076 3.789l-9.641-10.396v50.819h-5.562V16.217l-9.641 10.396z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M19.502 22.824L36 5.036l16.498 17.788l-4.076 3.789l-9.641-10.396v50.819h-5.562V16.217l-9.641 10.396z"/>`),
		g.Group(children),
	)
}