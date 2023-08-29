package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaOrbit0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaOrbit1" cx="208" cy="48.004" r="32.004"/><path id="galaOrbit2" d="M 226.27213,74.276705 A 112.01361,112.01361 0 0 1 195.13671,217.67055 112.01361,112.01361 0 0 1 48.774885,207.20976 112.01361,112.01361 0 0 1 38.349122,60.84543 112.01361,112.01361 0 0 1 181.75042,29.744337"/><path id="galaOrbit3" d="M 219.63599,191.62385 A 112.01361,112.01361 0 0 0 127.86254,144.01566 112.01361,112.01361 0 0 0 36.197622,191.83249"/></g>`),
		g.Group(children),
	)
}