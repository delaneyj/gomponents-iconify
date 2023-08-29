package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1248 0q119 0 203.5 84.5T1536 288v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960zM698 768q0-88-62-150t-150-62t-150 62t-62 150t62 150t150 62t150-62t62-150zm564 0q0-88-62-150t-150-62t-150 62t-62 150t62 150t150 62t150-62t62-150z"/>`),
		g.Group(children),
	)
}