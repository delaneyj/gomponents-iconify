package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43zm256 85q0-27-18.5-45.5T192 64t-45.5 18.5T128 128t18.5 45.5T192 192t45.5-18.5T256 128zM64 299v21h256v-21q0-20-23.5-36T244 240t-52-7t-52 7t-52.5 23T64 299z"/>`),
		g.Group(children),
	)
}