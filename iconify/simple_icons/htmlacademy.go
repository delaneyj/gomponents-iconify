package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Htmlacademy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0L2.524.994v17.368L12 24l9.476-5.638V.994L12.099.01L12 0zm8.236 17.657L12 22.557l-8.236-4.9v-7.119l8.2 4.881l.014.885l-5.626-3.349l-.008.86l5.648 3.394l.015.908l-5.647-3.36l-.008.86L12 19.01l5.703-3.412v-.862l-.008.004v-2.805l2.54-1.517v7.238zm-.006-8.162l-2.254 1.328l-1.04.613l-4.96-2.951l-.009.858l4.24 2.521l-.037.023l-.092.054l-.602.355l-3.5-2.083l-.009.859l2.763 1.643l-.652.436l-.015.01l-2.088-1.23l-.008.858l1.37.807l-1.395.837l-8.16-4.85l8.172-4.912v.001l8.276 4.823zm.006-.864l-8.28-4.882h-.002l-8.19 4.877V2.11L12 1.246l8.237.864v6.52z"/>`),
		g.Group(children),
	)
}