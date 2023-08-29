package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0zM289.966 323.877h620.068v405.835H289.966V323.877zm76.904 73.975v257.959h466.26V397.852H366.87zM212.256 766.626h775.488v68.628l-39.404 40.869H251.66l-39.404-40.869v-68.628zm315.82 22.485v62.183h143.848v-62.183H528.076z"/>`),
		g.Group(children),
	)
}