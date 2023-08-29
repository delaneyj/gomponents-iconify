package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<g fill="none"><path fill="#006A4E" d="M0 .75h300v185.5H0z"/><path fill="#FFCE00" d="M0 37.85h300v37.1H0zm0 74.2h300v37.1H0z"/><path fill="#D21034" d="M0 .75h111.248v111.3H0z"/><path fill="#FFF" d="M22.12 45.509L42.827 60.56l-7.91 24.354l20.707-15.052l20.707 15.052l-7.909-24.354l20.707-15.051H63.534l-7.91-24.354l-7.909 24.354z"/></g>`),
		g.Group(children),
	)
}