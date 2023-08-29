package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M54.574 12.771L35.5 40.333l-10.333-10v4.917l10.708 10.5l18.699-27.125z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M52.084 28.991A17.944 17.944 0 0 1 53.5 36c0 9.941-8.059 18-18 18s-18-8.059-18-18s8.059-18 18-18c3.267 0 6.33.87 8.971 2.392"/><path d="M54.574 12.771L35.5 40.333l-10.333-10v4.917l10.708 10.5l18.699-27.125z"/></g>`),
		g.Group(children),
	)
}