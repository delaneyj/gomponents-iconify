package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Navigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 1024q-40 0-96-64L576 832q-10-6-35-22t-36-22.5t-28-13t-29-6.5q-16 0-35 8.5T361.5 806T320 832L160 960q-3 3-14 13.5t-15.5 15t-15 13t-17.5 12t-16.5 7T64 1024q-27 0-45.5-19T0 960L384 64q0-26 18.5-45t45-19T493 19t19 45l384 896q0 27-19 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}