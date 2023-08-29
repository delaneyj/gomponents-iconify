package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1472 1216H192q-24 0-44 38.5t-20 89.5t20 89.5t44 38.5h1280q24 0 44-38.5t20-89.5t-20-89.5t-44-38.5zM128 128q-53 0-90.5 37.5T0 256t37.5 90.5T128 384l80 703h1248l80-703q53 0 90.5-37.5T1664 256t-37.5-90.5T1536 128t-90.5 37.5T1408 256q0 56 41 94q-153 183-207 236.5t-90 53.5q-34 0-83-66.5T899 301q28-17 44.5-46t16.5-63q0-53-37.5-90.5T832 64t-90.5 37.5T704 192q0 34 16.5 63t44.5 46Q644 507 595 573.5T512 640q-36 0-90-53.5T215 350q41-38 41-94q0-53-37.5-90.5T128 128z"/>`),
		g.Group(children),
	)
}