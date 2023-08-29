package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeatherShower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 18a1 1 0 1 1 0-2c1.654 0 3-1.346 3-3s-1.346-3-3-3c-.243 0-.5.041-.81.13l-1.075.307l-.185-1.103A3.98 3.98 0 0 0 11 6a4.004 4.004 0 0 0-3.918 4.811l.244 1.199l-1.42-.016C4.897 12 4 12.897 4 14s.897 2 2 2a1 1 0 1 1 0 2c-2.206 0-4-1.794-4-4a4.007 4.007 0 0 1 3.001-3.874L5 10c0-3.309 2.691-6 6-6a5.969 5.969 0 0 1 5.65 4.015C19.589 7.771 22 10.128 22 13c0 2.757-2.243 5-5 5zm-6.5 0l1-3l1 3a1.053 1.053 0 1 1-2 0zm3 2l1-3l1 3a1.053 1.053 0 1 1-2 0zm-6 0l1-3l1 3a1.053 1.053 0 1 1-2 0z"/>`),
		g.Group(children),
	)
}