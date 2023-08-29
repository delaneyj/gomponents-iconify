package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swipedown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.297 1024h-384q-21 0-51.5-34.5t-53.5-81t-24-77.5q0-24-16-81.5t-31.5-131.5t-15.5-138v-32q0-27 18.5-45.5t45.5-18.5q22 0 39 13t22 34q3-28 3-47q0-59 2-113t8-122t19.5-108.5t34.5-40.5q45 0 54.5 45t9.5 211q0 37 6.5 64t16 39t19 18t15.5 7h7q0-27 18.5-45.5t45.5-18.5q64 0 64 128q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v64q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v288q0 46-7.5 83t-18 59.5t-25 39.5t-27 24.5t-25.5 12t-17.5 5t-7.5.5zm-724-518q-5 6-12 6t-12-6l-145-163q-8-9 5-23h120V32q0-14 9-23t22.5-9t23 9t9.5 23v288h120q13 14 4 23z"/>`),
		g.Group(children),
	)
}