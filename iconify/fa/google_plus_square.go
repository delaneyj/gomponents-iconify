package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlusSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M917 777q0-26-6-64H549v132h217q-3 24-16.5 50T712 948t-66.5 44.5T549 1010q-99 0-169-71t-70-171t70-171t169-71q92 0 153 59l104-101Q698 384 549 384q-160 0-272 112.5T165 768t112 271.5T549 1152q165 0 266.5-105T917 777zm345 46h109V713h-109V603h-110v110h-110v110h110v110h110V823zm274-535v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}