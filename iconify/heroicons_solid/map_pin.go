package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m9.69 18.933l.003.001C9.89 19.02 10 19 10 19s.11.02.308-.066l.002-.001l.006-.003l.018-.008a5.741 5.741 0 0 0 .281-.14c.186-.096.446-.24.757-.433c.62-.384 1.445-.966 2.274-1.765C15.302 14.988 17 12.493 17 9A7 7 0 1 0 3 9c0 3.492 1.698 5.988 3.355 7.584a13.731 13.731 0 0 0 2.273 1.765a11.842 11.842 0 0 0 .976.544l.062.029l.018.008l.006.003ZM10 11.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}