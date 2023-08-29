package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarWithLeftHalfBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="m37.06 11l-9.208 15.87L9.5 29.544l13.285 12.94l-3.128 18.28L37.06 52.5V11Z"/><path fill="#FCEA2B" d="m35 11.5l9.648 15.87L63 30.044l-13.285 12.94l3.128 18.28L35 52.5v-41Z"/><path fill="#3F3F3F" d="m37.06 11l-9.208 15.87L9.5 29.544l13.285 12.94l-3.128 18.28L37.06 52.5V11Z"/><g stroke="#000" stroke-linecap="round" stroke-width="2"><path fill="none" stroke-linejoin="round" stroke-miterlimit="10" d="M35.993 10.736L27.79 27.37L9.44 30.044l13.285 12.94l-3.128 18.28l16.412-8.636l16.419 8.623l-3.142-18.277l13.276-12.95l-18.354-2.66l-8.214-16.628Z"/><path d="M36 12v37"/></g>`),
		g.Group(children),
	)
}