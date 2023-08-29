package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sensor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.826 3 3 8.826 3 16s5.826 13 13 13s13-5.826 13-13S23.174 3 16 3zm0 1c6.633 0 12 5.367 12 12s-5.367 12-12 12S4 22.633 4 16S9.367 4 16 4zm0 2C10.49 6 6 10.49 6 16s4.49 10 10 10s10-4.49 10-10S21.51 6 16 6zm0 2c4.43 0 8 3.57 8 8s-3.57 8-8 8s-8-3.57-8-8s3.57-8 8-8zm0 2c-3.302 0-6 2.698-6 6s2.698 6 6 6s6-2.698 6-6s-2.698-6-6-6zm0 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4c-2.22 0-4-1.78-4-4c0-2.22 1.78-4 4-4zm0 2a2 2 0 1 0-.001 3.999A2 2 0 0 0 16 14z"/>`),
		g.Group(children),
	)
}