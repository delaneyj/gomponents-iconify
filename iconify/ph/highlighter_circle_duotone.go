package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlighterCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M195.88 195.88A95.7 95.7 0 0 1 168 215.29V152a8 8 0 0 0-8-8h-8V72l-48 24v48h-8a8 8 0 0 0-8 8v63.29a96 96 0 1 1 107.88-19.41Z" opacity=".2"/><path d="M201.54 54.46A104 104 0 0 0 54.46 201.54A104 104 0 0 0 201.54 54.46ZM96 210v-58h64v58a88.33 88.33 0 0 1-64 0Zm48-74h-32v-35.06l32-16Zm46.22 54.22A88.09 88.09 0 0 1 176 201.77V152a16 16 0 0 0-16-16V72a8 8 0 0 0-11.58-7.16l-48 24A8 8 0 0 0 96 96v40a16 16 0 0 0-16 16v49.77a88 88 0 1 1 110.22-11.55Z"/></g>`),
		g.Group(children),
	)
}