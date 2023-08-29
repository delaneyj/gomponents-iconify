package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#FFF" d="M.5.5h300v200H.5z"/><path fill="#01411C" d="M75.5.5h225v200h-225z"/><circle cx="188" cy="100.5" r="60" fill="#FFF"/><circle cx="203.337" cy="86.868" r="55" fill="#01411C"/><path fill="#FFF" d="m209.879 55.603l25.274 28.433l-37.16-8.147l34.852-15.251l-19.232 32.824z"/></g>`),
		g.Group(children),
	)
}