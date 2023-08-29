package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Like(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 640h-64q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19q27 0 45.5 18.5t18.5 45t-18.5 45.5t-45.5 19h-64q27 0 45.5 18.5T896 960t-18.5 45.5T832 1024H320q-32 0-128-32T64 960q-28 0-46-16.5T0 896V640q0-35 18.5-79T60 513q1-1 4-1q54 0 105-14.5t85.5-34t64.5-44t44-40.5t21-27q20-30 56.5-66.5t64.5-62t49.5-60T576 96V64q0-27 19-45.5T640 0t45 18.5T704 64v128q0 43-10 74.5T672 312t-22 32.5t-10 39.5q0 18 10 31t24 19t28 9.5t24 3.5l10 1q117 0 202.5 39.5T1024 576q0 26-18.5 45T960 640z"/>`),
		g.Group(children),
	)
}