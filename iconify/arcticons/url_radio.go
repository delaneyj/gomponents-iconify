package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UrlRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.005 11.676h35.986a1.509 1.509 0 0 1 1.509 1.508V34.82a1.506 1.506 0 0 1-1.506 1.505H6.01a1.509 1.509 0 0 1-1.51-1.509V13.18a1.506 1.506 0 0 1 1.505-1.505Z"/><ellipse cx="15.278" cy="23.787" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="7.355" ry="7.339"/><circle cx="27.985" cy="30.376" r="1.417" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.349" cy="30.376" r="1.417" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.723" cy="30.375" r="1.417" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.178 17.247h10.348a1.614 1.614 0 0 1 1.614 1.614v5.986a1.61 1.61 0 0 1-1.61 1.61H28.178a1.61 1.61 0 0 1-1.61-1.61v-5.99a1.61 1.61 0 0 1 1.61-1.61Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.54 20.187v3.331h1.669m-5.836 0v-3.331h1.085a1.124 1.124 0 1 1 0 2.248h-1.085m1.127-.041l1.085 1.124m-6.087-3.332v2.207a1.085 1.085 0 1 0 2.17 0v-2.207"/>`),
		g.Group(children),
	)
}