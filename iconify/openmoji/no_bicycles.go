package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBicycles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26" fill="#fff"/><path fill="#ea5a47" d="M36 7C19.98 7 7 19.98 7 36s12.98 29 29 29s29-12.98 29-29S52.02 7 36 7zM10.79 36.27c0-5.075 3.417-12.69 6.202-16.49l35.2 35.27c-3.805 2.784-10.93 5.904-16.01 5.904c-12.7 0-25.39-11.98-25.39-24.68zm44.87 15.18L20.6 16.32c3.805-2.784 10.77-5.441 15.84-5.441c12.7 0 24.68 12.25 24.68 24.95c0 5.075-2.686 11.81-5.47 15.62z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><circle cx="36" cy="36" r="30" stroke-width="2"/><path stroke-width="2" d="M55.66 51.44A24.99 24.99 0 0 0 20.6 16.31zM16.99 19.77a24.99 24.99 0 0 0 35.2 35.27z"/><path stroke-width="2.727" d="M42 25.7h6" transform="matrix(.7371 0 0 .7297 8.857 6.124)"/><path stroke-width="2" d="m24.41 39.69l1.754-5.673M41.42 28.01l4.703 9.462M36.74 28.01h4.683m2.217-3.13l-3.986 6.085m2.816 2.955a7.439 7.439 0 0 1 9.826.307a7.439 7.439 0 0 1 .94 9.786M31.93 39.69a7.439 7.439 0 0 1-7.429 7.439a7.439 7.439 0 0 1-7.449-7.419a7.439 7.439 0 0 1 7.409-7.459"/></g>`),
		g.Group(children),
	)
}