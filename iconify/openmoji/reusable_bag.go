package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReusableBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="32.615" height="38.612" x="19.693" y="20.941" fill="#a57939" rx="6.783"/><path fill="#b1cc33" d="M33.427 38.127c1.806-2.92 5.868-3.617 5.868-3.617s1.118 3.996-.67 6.839s-5.869 3.616-5.869 3.616s-1.136-3.919.67-6.838Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M33.645 38.054c1.806-2.92 5.868-3.616 5.868-3.616s1.118 3.996-.67 6.838s-5.869 3.617-5.869 3.617s-1.136-3.919.67-6.839ZM32.1 46.457l.807-1.658"/><g stroke-miterlimit="10"><path stroke-linecap="round" d="M28.157 25.663c0-7.445 3.416-13.48 7.63-13.48s7.631 6.035 7.631 13.48"/><rect width="32.615" height="38.612" x="19.693" y="20.941" rx="6.783"/></g></g>`),
		g.Group(children),
	)
}