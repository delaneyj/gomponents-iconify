package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1755 453q37 38 37 90.5t-37 90.5l-401 400l150 150l-160 160q-163 163-389.5 186.5T543 1430l-362 362H0v-181l362-362q-124-185-100.5-411.5T448 448l160-160l150 150l400-401q38-37 91-37t90 37t37 90.5t-37 90.5L939 619l234 234l401-400q38-37 91-37t90 37z"/>`),
		g.Group(children),
	)
}