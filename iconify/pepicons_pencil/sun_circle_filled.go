package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilSunCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" transform="translate(3 3)"><path fill-rule="evenodd" d="M10 13.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-6a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z" clip-rule="evenodd"/><rect width="1" height="3" x="1" y="10.5" rx=".5" transform="rotate(-90 1 10.5)"/><rect width="1" height="3" x="16" y="10.5" rx=".5" transform="rotate(-90 16 10.5)"/><rect width="1" height="3" x="14" y="14.707" rx=".5" transform="rotate(-45 14 14.707)"/><rect width="1" height="3" x="3" y="3.707" rx=".5" transform="rotate(-45 3 3.707)"/><rect width="1" height="3" x="9.5" y="16" rx=".5"/><rect width="1" height="3" x="9.5" y="1" rx=".5"/><rect width="1" height="3" x="16.121" y="3" rx=".5" transform="rotate(45 16.121 3)"/><rect width="1" height="3" x="5.121" y="14" rx=".5" transform="rotate(45 5.121 14)"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilSunCircleFilled0)"/></g>`),
		g.Group(children),
	)
}