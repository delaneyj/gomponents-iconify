package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M35.993 10.736L27.79 27.37L9.44 30.044l13.285 12.94l-3.128 18.28l16.412-8.636l16.419 8.623l-3.142-18.277l13.276-12.95l-18.354-2.66l-8.214-16.628Z"/><path fill="#3F3F3F" d="M35.993 10.736L27.79 27.37L9.44 30.044l13.285 12.94l-3.128 18.28l16.412-8.636l16.419 8.623l-3.142-18.277l13.276-12.95l-18.354-2.66l-8.214-16.628Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M35.993 10.736L27.79 27.37L9.44 30.044l13.285 12.94l-3.128 18.28l16.412-8.636l16.419 8.623l-3.142-18.277l13.276-12.95l-18.354-2.66l-8.214-16.628Z"/>`),
		g.Group(children),
	)
}