package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StadtradHamburg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="30.686" cy="37.715" r="3.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.314" cy="37.715" r="3.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.616 31.425l-4.89 6.285h-6.41l4.908-6.285h6.392"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.686 37.715l-2.843-8.621H25.75m-4 .342l1.974 8.276m-3.252-8.276h2.662"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.682 26.997a59.416 59.416 0 0 1-3.78 16.108q17.102.79 34.196 0a58.266 58.266 0 0 1-3.78-16.108m-28.616 0v-2.34h30.596v2.34Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.682 24.657V13.86h6.389v10.798m3.329-.001V12.06m7.2 0v12.597M24 8.46V4.5m-1.8 1.8h3.6m5.13 18.357V13.86h6.388v10.798M20.4 12.06a3.6 3.6 0 1 1 7.2 0m.009.449h1.342m-8.551 0h-1.35m-5.203-.378l-.689-1.61l-1.742.17l1.05-1.401l-1.02-1.424l1.738.208l.723-1.594l.69 1.61l1.742-.17l-1.049 1.4l1.019 1.425l-1.739-.208Zm20.248 0l-.69-1.61l-1.742.17l1.05-1.401l-1.02-1.424l1.739.208l.723-1.594l.689 1.61l1.742-.17l-1.049 1.4l1.019 1.425l-1.738-.208Z"/>`),
		g.Group(children),
	)
}