package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basketball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M685 641L557 512l338-339q62 70 95.5 156.5T1024 512q-96 0-182.5 34T685 641zM383 339q62-70 95.5-156.5T512 0q96 0 182.5 34T851 129L512 467zM174 129Q292 24 448 5q-2 164-110 288zm119 209Q169 446 4 448q20-156 125-274zm174 174L129 851q-61-70-95-156.5T0 512q96 0 182.5-33.5T339 383zm174 173q-62 70-95.5 156.5T512 1024q-96 0-182.5-33.5T173 895l339-338zm209 210q-118 105-274 125q2-164 110-289zm170-319q-20 156-125 274L731 686q124-108 289-110z"/>`),
		g.Group(children),
	)
}