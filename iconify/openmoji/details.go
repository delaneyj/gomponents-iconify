package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Details(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M51.065 43.916V10.979H18.937v49.76h16.062"/><circle cx="43.167" cy="52.167" r="11.129" fill="#D0CFCE"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M51.065 43.916V10.979H18.937v49.76h16.062M23.667 16.02H32m-8.333 8.48h22m-22 3.771h22m-22 3.77h22m-22 3.771h22m-22 3.771h22M36.38 43.35H23.67m9.58 3.77h-9.58m8.45 3.77h-8.45"/><circle cx="43.167" cy="52.167" r="11.129"/><path d="m51.065 60.739l5.467 5.467M33.439 48.453h13.436m-14.703 5.386h14.703"/></g>`),
		g.Group(children),
	)
}