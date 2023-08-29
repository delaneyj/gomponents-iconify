package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuppetIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FFAE1A" d="M256 256V138.204h-90.155l-48.049-48.307V0H0v117.796h90.155l48.307 48.307v62.256l-48.307 48.307H0v117.796h117.796v-89.897l48.307-48.307H256V256ZM39.265 39.265h39.266v39.266H39.265V39.265Zm39.266 315.673H39.265v-39.265h39.266v39.265Z"/>`),
		g.Group(children),
	)
}