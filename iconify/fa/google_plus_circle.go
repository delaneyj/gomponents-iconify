package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M917 777q0-33-6-64H549v132h217q-12 76-74.5 120.5T549 1010q-99 0-169-71.5T310 768t70-170.5T549 526q93 0 153 59l104-101Q698 384 549 384q-160 0-272 112.5T165 768t112 271.5T549 1152q165 0 266.5-105T917 777zm345 46h109V713h-109V603h-110v110h-110v110h110v110h110V823zm274-55q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}