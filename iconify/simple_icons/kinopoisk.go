package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kinopoisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.273 16.103c-2.266 6.226-9.15 9.436-15.376 7.17C1.67 21.007-1.54 14.123.727 7.897c2.266-6.226 9.15-9.436 15.376-7.17c6.226 2.266 9.436 9.15 7.17 15.376zm-9.864-3.59a1.5 1.5 0 1 0-2.818-1.026a1.5 1.5 0 0 0 2.818 1.026zm-4.357 3.202a3 3 0 1 0-5.636-2.052a3 3 0 0 0 5.636 2.052zM12.13 7.26a3 3 0 1 0-5.636-2.052A3 3 0 0 0 12.13 7.26zm8.456 3.077a3 3 0 1 0-5.637-2.052a3 3 0 0 0 5.636 2.052zm-3.078 8.455a3 3 0 1 0-5.637-2.051a3 3 0 0 0 5.637 2.05Z"/>`),
		g.Group(children),
	)
}