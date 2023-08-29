package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digikeyelectronics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.246.221A11.786 11.786 0 0 1 23.89 10.418c.912 6.593-3.944 12.711-10.558 13.297c-.454.04-.912.063-1.369.064l-10.705.003v-3.749H0V3.987h1.222V.218l11.024.003zM17.9 19.423l-8.26-7.422l8.25-7.422h-6.938L5.615 9.361V4.598H.56v14.803h5.105v-4.724l5.289 4.746H17.9z"/>`),
		g.Group(children),
	)
}