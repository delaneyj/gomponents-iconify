package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarWithRightHalfBlack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="m35 11l9.209 15.87l18.352 2.674l-13.285 12.94l3.128 18.28L35 52.5V11Z"/><path fill="#FCEA2B" d="m37 11.5l-9.209 15.87L9.44 30.044l13.285 12.94l-3.128 18.28L37 52V11.5Z"/><path fill="#3F3F3F" d="m35 11l9.209 15.87l18.352 2.674l-13.285 12.94l3.128 18.28L35 52.5V11Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" stroke-miterlimit="10" d="M35.993 10.736L27.79 27.37L9.44 30.044l13.285 12.94l-3.128 18.28l16.412-8.636l16.419 8.623l-3.142-18.277l13.276-12.95l-18.354-2.66l-8.214-16.628Z"/><path d="M36 11.5V49"/></g>`),
		g.Group(children),
	)
}