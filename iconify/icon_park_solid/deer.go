package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDeer0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-width="4" d="M36 27c0 7.217-5.373 17-12 17s-12-9.783-12-17c0-7.216 1.5-11 12-11s12 3.784 12 11Z"/><ellipse fill="#fff" stroke="#fff" stroke-width="4" rx="5" ry="3.5" transform="scale(1 -1) rotate(45 40.625 38.327)"/><ellipse cx="9" cy="17.5" fill="#fff" stroke="#fff" stroke-width="4" rx="5" ry="3.5" transform="rotate(45 9 17.5)"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M12 4c0 6.627 5.373 12 12 12s12-5.373 12-12"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M18 7c0 4.97 2.686 9 6 9s6-4.03 6-9"/><circle cx="20" cy="26" r="2" fill="#000"/><circle cx="24" cy="34" r="2" fill="#000"/><circle cx="28" cy="26" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDeer0)"/>`),
		g.Group(children),
	)
}