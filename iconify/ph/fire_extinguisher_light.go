package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisherLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M217.72 50.25L152.21 30.6l34.47-17.23a6 6 0 1 0-5.36-10.74L134.51 26A78.07 78.07 0 0 0 58 104v104a6 6 0 0 0 12 0v-34h20v58a14 14 0 0 0 14 14h64a14 14 0 0 0 14-14V104a46.07 46.07 0 0 0-40-45.6V40.07l72.27 21.68A6.14 6.14 0 0 0 216 62a6 6 0 0 0 1.72-11.75ZM70 162v-58a66.07 66.07 0 0 1 60-65.71V58.4A46.07 46.07 0 0 0 90 104v58Zm98 72h-64a2 2 0 0 1-2-2v-58h68v58a2 2 0 0 1-2 2Zm2-130v58h-68v-58a34 34 0 0 1 68 0Z"/>`),
		g.Group(children),
	)
}