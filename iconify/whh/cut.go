package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 833q-37 0-68-19.5T781 761l-38-17l17 38q32 16 52 47t20 68q0 53-37.5 90.5T704 1025t-90.5-37.5T576 897q0-34 17-63t45-46l-50-113l-76-34Q224 480 102 389.5T4 258l508 255L257 4q41-24 130.5 97.5T640 513l34 76l113 50q17-28 46-45t63-17q53 0 90.5 37.5T1024 705t-37.5 90.5T896 833zm-256 64.5q0 26.5 19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19t-45.5 19t-19 45.5zM896.5 641q-26.5 0-45.5 19t-19 45.5t19 45t45.5 18.5t45-18.5t18.5-45t-18.5-45.5t-45-19z"/>`),
		g.Group(children),
	)
}