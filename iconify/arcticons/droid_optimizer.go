package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DroidOptimizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.737 12.536L14.524 13.6l-3.686-1.09s-.978.041-1.276.342s-1.967 2.8-1.967 2.8s-.522 1.175-.106 1.74s2.948 3.042 2.948 3.042l-.075 1.626l-3.9 1.832a1.385 1.385 0 0 0-.594 1.094a12.423 12.423 0 0 0 .566 3.514a2.526 2.526 0 0 0 .771.861a3.921 3.921 0 0 0 1.542.506c.669.001 2.78.083 2.78.083l.782 1.292l-1.612 3.208s-.42 1.308-.113 1.704s2.816 2.498 2.816 2.498a1.435 1.435 0 0 0 1.808.19c1.048-.634 3.317-2.24 3.317-2.24l1.386.63l.835 3.336s.286 1.123 1.043 1.271s4.12.08 4.12.08a2.183 2.183 0 0 0 1.228-1.328c.24-.938.881-3.36.881-3.36l1.437-.564c1.548 1.351 4.267 2.528 4.65 2.231c.728-.44 2.722-2.068 2.722-2.068a1.562 1.562 0 0 0 .405-2.168c-.782-1.405-1.617-3.205-1.617-3.205l.842-1.369l3.824-.226c2.368-.04 1.758-3.035 1.825-5.127c-.217-1.218-2.569-1.757-4.27-2.57l-.231-1.524l3.132-3.211a2.08 2.08 0 0 0-.325-1.695c-.59-.947-1.333-2.384-1.333-2.384s-.341-.895-2.063-.563a34.726 34.726 0 0 0-3.402.864l-1.266-.88l.347-3.898a1.338 1.338 0 0 0-.768-1.565c-1.054-.424-3.157-1.156-3.157-1.156a1.774 1.774 0 0 0-1.848.65c-.602.959-2.04 3.094-2.04 3.094l-1.58.015l-1.97-2.993a1.618 1.618 0 0 0-2.072-.757a33.136 33.136 0 0 0-3.142 1.145s-.782.067-.68 1.794s.3 3.435.3 3.435Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="9.163" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="5.872" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="12.054" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}