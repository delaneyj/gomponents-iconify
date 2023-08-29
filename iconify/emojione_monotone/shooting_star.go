package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShootingStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m7 30.623l13.858 10.002L13.71 62l18.706-13.206l17.656 12.467L57 2L7 30.623M55.591 4.095l-.018.135l-13.4 24.108h-3.3l-1.162-3.48L55.488 4.155l.103-.06M42.679 50.701l-10.263-7.246l-10.263 7.246l3.921-11.726l-10.264-7.244h12.685l3.922-11.725l3.921 11.725h12.685L38.76 38.976l3.919 11.725M36.424 21.01l-1.557-4.655L55.155 4.353L36.424 21.01m10.18 7.329l8.922-23.744l-3.099 23.744h-5.823"/>`),
		g.Group(children),
	)
}