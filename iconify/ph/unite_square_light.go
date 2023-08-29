package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UniteSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 90h-50V40a6 6 0 0 0-6-6H40a6 6 0 0 0-6 6v120a6 6 0 0 0 6 6h50v50a6 6 0 0 0 6 6h120a6 6 0 0 0 6-6V96a6 6 0 0 0-6-6Zm-61.52 120L46 101.52v-47L201.52 210Zm-100-164h47L210 154.48v47ZM210 137.52L174.48 102H210Zm-56-56L118.48 46H154Zm-108 37L81.52 154H46Zm56 56L137.52 210H102Z"/>`),
		g.Group(children),
	)
}