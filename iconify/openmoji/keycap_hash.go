package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapHash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.208h48v47.834h-48z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.125 11.958h48v48h-48z"/><path d="M33.453 27.616L30.94 44.634"/><path d="M40.81 27.616l-2.513 17.018"/><path d="M28.066 32.357l16.717-.017"/><path d="M26.967 39.847l16.717-.017"/></g>`),
		g.Group(children),
	)
}