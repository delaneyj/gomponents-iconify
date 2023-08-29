package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTrafficLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M47 15.82v36c0 3.85-3.15 7-7 7h-8c-3.85 0-7-3.15-7-7v-36c0-3.85 3.15-7 7-7h8c3.85 0 7 3.15 7 7z"/><circle cx="36" cy="48.292" r="5" fill="#b1cc33"/><circle cx="36" cy="34.292" r="5" fill="#f4aa41"/><circle cx="36" cy="19.292" r="5" fill="#ea5a47"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M47 16v36c0 3.85-3.15 7-7 7h-8c-3.85 0-7-3.15-7-7V16c0-3.85 3.15-7 7-7h8c3.85 0 7 3.15 7 7z"/><circle cx="36" cy="48.472" r="5"/><circle cx="36" cy="34.472" r="5"/><circle cx="36" cy="19.472" r="5"/></g>`),
		g.Group(children),
	)
}