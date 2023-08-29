package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidNineteenVirusFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.25 19.651a3.714 3.714 0 1 0 0-7.428a3.714 3.714 0 0 0 0 7.428ZM6.631 9.437h1.238m-.619 0v2.786m4.158-1.32l.876.875m-.438-.437l-1.97 1.969m3.874 2.008v1.238m0-.619h-2.786m1.32 4.158l-.876.876m.438-.438l-1.97-1.97m-2.007 3.874H6.631m.619 0v-2.786m-4.158 1.32l-.876-.876m.438.438l1.97-1.97M.75 16.556v-1.238m0 .619h2.786m-1.32-4.159l.876-.875m-.438.438l1.97 1.969M18.5 9.027a2.714 2.714 0 1 0 0-5.428a2.714 2.714 0 0 0 0 5.428Zm-.452-7.464h.904m-.452 0v2.036m3.039-.964l.64.639m-.32-.32l-1.44 1.44m2.831 1.467v.905m0-.453h-2.036m.965 3.039l-.64.64m.32-.32l-1.44-1.44m-1.467 2.831h-.904m.452 0V9.027m-3.039.965l-.64-.64m.32.32l1.44-1.44M13.75 6.766v-.905m0 .452h2.036m-.965-3.039l.64-.639m-.32.319l1.44 1.44"/>`),
		g.Group(children),
	)
}