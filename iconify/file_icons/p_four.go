package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472.667 417.35v-95.248h38.952v-54.667h-38.952V61.102H406l-161.333 207v54H410v95.247H92V278.102c121.35 3.198 167.824-45.821 170.788-114.586c3.509-81.401-42.069-102.414-106.789-102.414h-137v356.247H0v33.55h512v-33.55h-39.333zM303.528 267.434L410 133.102v134.333H303.528zM92 111.102s95.613-11.625 95.613 45.648c0 81.36-95.613 71.608-95.613 71.608V111.102z"/>`),
		g.Group(children),
	)
}