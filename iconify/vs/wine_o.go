package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M547 2190q227 2 394-35q41-9 72.5-40.5t48-74t24.5-85.5t8-84v-589q0-156-70-349q-17-46-64-112.5T868 700t-80.5-107.5T752 514V139q0-71-12.5-105T692 0H402q-36 0-48 33.5T342 139v375q0 25-35.5 78.5T226 700t-92 120.5T70 933Q0 1114 0 1282v589q0 41 8 83.5t24.5 85.5t47.5 74.5t72 40.5q164 37 395 35zM465 138h162v393q0 47 29 107.5T721.5 751T801 872.5T864 997q62 164 62 280v581q0 75-20 112t-83 48q-132 23-276 23t-276-23q-64-11-84-48t-20-112v-581q0-116 62-280q20-55 62.5-124.5T371 751t65.5-112.5T465 531V138z"/>`),
		g.Group(children),
	)
}