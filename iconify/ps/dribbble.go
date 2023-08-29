package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dribbble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2zm141 119q37 46 38 103q-39-9-73-9q-25 0-52 5q-4-10-12-28q60-27 99-71zM232 52q62 0 112 40q-30 37-88 63q-26-50-62-98q20-5 38-5zm-79 19q36 45 64 98q-69 20-142 20h-3q-10 0-14-1q19-80 95-117zM52 232v-3q5 1 20 1q90 0 163-24q2 4 5.5 12t5.5 12q-57 19-102 60q-32 29-49 58q-43-51-43-116zm180 180q-60 0-107-36q23-33 42-51q40-39 93-57q21 58 36 132q-30 12-64 12zm102-32q-11-55-33-121q17-3 36-3h1q32 0 70 9q-12 71-74 115z"/>`),
		g.Group(children),
	)
}