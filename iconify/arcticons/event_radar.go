package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventRadar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.985 13.23c-.014.312-.047.619-.047.934a19.97 19.97 0 0 0 32.124 15.87Zm15.905 8.332l1.32-2.303"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.31 33.83S24 37.99 27.2 43.5H10.07a89.014 89.014 0 0 0 5.071-13.969M39.03 15.654a12.813 12.813 0 0 0-18.233-9.9"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.517 16.11a9.269 9.269 0 0 0-13.19-7.161"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.004 16.567a5.725 5.725 0 0 0-8.146-4.423"/><circle cx="25.744" cy="18.328" r="1.077" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}