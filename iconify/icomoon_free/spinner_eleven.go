package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 6h-6l2.243-2.243C11.11 2.624 9.603 2 8 2s-3.109.624-4.243 1.757C2.624 4.89 2 6.397 2 8s.624 3.109 1.757 4.243A5.961 5.961 0 0 0 8 14a5.963 5.963 0 0 0 4.516-2.049l1.505 1.317a8 8 0 1 1-.364-10.924L16 0v6z"/>`),
		g.Group(children),
	)
}