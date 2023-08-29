package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" fill-rule="evenodd" d="M0 0h512v512H0Z"/><g fill-rule="evenodd" transform="rotate(-56.3 367.2 -111.2) scale(9.375)"><g id="flagKr1x10"><path id="flagKr1x11" fill="#000" d="M-6-26H6v2H-6Zm0 3H6v2H-6Zm0 3H6v2H-6Z"/><use width="100%" height="100%" y="44" href="#flagKr1x11"/></g><path stroke="#fff" d="M0 17v10"/><path fill="#cd2e3a" d="M0-12a12 12 0 0 1 0 24Z"/><path fill="#0047a0" d="M0-12a12 12 0 0 0 0 24A6 6 0 0 0 0 0Z"/><circle cy="-6" r="6" fill="#cd2e3a"/></g><g fill-rule="evenodd" transform="rotate(-123.7 196.5 59.5) scale(9.375)"><use width="100%" height="100%" href="#flagKr1x10"/><path stroke="#fff" d="M0-23.5v3M0 17v3.5m0 3v3"/></g>`),
		g.Group(children),
	)
}