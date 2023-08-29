package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="49.5" cy="51.5" r="3" fill="#d0cfce"/><circle cx="49.5" cy="51.5" r="1.5" fill="#9b9b9a"/><path fill="none" stroke="#9b9b9a" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M30.1 13c-.576.49-2.052.652-1.461 5.457C29.144 22.56 33 38.5 36.5 38.5s7.163-16.959 7.86-20.057c.58-2.583.13-3.095-1.36-4.943"/><g stroke="#000" stroke-miterlimit="10"><path stroke-linecap="round" d="M41.203 9.905c1.155.667 1.883 1.087 1.279 2.133s-1.332.626-2.487-.041s-2.092-1.208-1.488-2.254s1.54-.505 2.696.162Z"/><path fill="none" stroke-linecap="round" stroke-width="2" d="M31.5 11.5c-2 1-5 2-3.861 6.957c.958 4.17 3.543 21.043 8.861 21.043c5.251 0 8-18 8.86-21.057c.854-3.038-1.86-5.943-3.86-6.943"/><path fill="none" stroke-width="2" d="M47.5 49.5c-5-6-5 8-10 12c-11.713 9.37-21-9-7-11c5.94-.849 6-11 6-11"/><circle cx="49.5" cy="51.5" r="3" fill="none"/><path stroke-linecap="round" d="M32.331 9.967c1.11-.404 1.81-.658 2.175.347s-.333 1.26-1.443 1.663s-2.01.732-2.376-.273s.534-1.333 1.644-1.737Z"/></g>`),
		g.Group(children),
	)
}