package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loudspeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="59.383" cy="36.38" r="5" fill="#3F3F3F" stroke="#3F3F3F" stroke-miterlimit="10" stroke-width="2"/><path fill="#d0cfce" d="M57.739 11.375L41.303 24.721H12.59a3.01 3.01 0 0 0-3 3v16.692a3.01 3.01 0 0 0 3 3h28.712L57.74 60.541l-.002-49.166zM18.751 25.333v21.25"/><path fill="#9b9b9a" d="m18.687 24.511l-7.535.437l-1.529 1.911l-.055 18.183l2.13 2.457l6.825-.055zm9.063 23.235V60.04c0 1.509-1.117 2.732-2.496 2.732h-4.992c-1.379 0-2.497-1.223-2.497-2.732V47.746"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M57.739 11.375L41.303 24.721H12.59a3.01 3.01 0 0 0-3 3v16.692a3.01 3.01 0 0 0 3 3h28.712L57.74 60.541l-.002-49.166zM18.751 25.333v21.25"/><path d="M26.75 51.445v9a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-9m42.864-19.538a5 5 0 0 1-.002 8.966"/></g>`),
		g.Group(children),
	)
}