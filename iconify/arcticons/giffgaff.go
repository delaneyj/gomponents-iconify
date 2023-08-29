package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Giffgaff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="17.57" cy="12.56" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 22.64v-8.58a1.82 1.82 0 0 1 1.83-1.82h0a2.09 2.09 0 0 1 1.83.76m-7.21 2.75h5.8m-3.39-3.47a3.65 3.65 0 0 0-.59 0h0a1.81 1.81 0 0 0-1.82 1.82v8.58h-2.88v-6.93H14.7v7.8a2.59 2.59 0 0 1-2.6 2.6h0a2.59 2.59 0 0 1-1.84-.76"/><rect width="5.2" height="6.89" x="9.5" y="15.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.6" transform="rotate(180 12.1 19.195)"/><rect width="5.2" height="6.89" x="15.25" y="25.94" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.6" transform="rotate(-180 17.845 29.39)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.84 32.84v-8.59a1.82 1.82 0 0 1 1.82-1.82h0a2.13 2.13 0 0 1 1.84.76m-7.25 2.75h5.81m-16.61 0v7.81a2.61 2.61 0 0 1-2.6 2.6h0a2.59 2.59 0 0 1-1.85-.76m12.38-5.35a2.6 2.6 0 0 1-2.6 2.6h0a2.61 2.61 0 0 1-2.6-2.6v-1.69a2.61 2.61 0 0 1 2.6-2.61h0a2.6 2.6 0 0 1 2.6 2.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.66 22.48a3.42 3.42 0 0 0-.59-.05h0a1.83 1.83 0 0 0-1.82 1.82v8.59h-2.87v-6.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}