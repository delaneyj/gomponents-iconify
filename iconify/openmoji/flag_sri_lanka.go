package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSriLanka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 17h11v38H5z"/><path fill="#e27022" d="M16 17h10v38H16z"/><path fill="none" stroke="#f1b31c" stroke-miterlimit="10" stroke-width="2" d="M26 18v36M7 19h58v34H7z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M50 45h2l1-2l-3-6m-9 8h2l1-2l-3-6m-4-11a7.07 7.07 0 0 0-1 4v12"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M44 41c3 0 3-2 7-2l5 2c-1-2 2-3 1-6H42s-3 2-3 5h3Z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M54 45h2l1-2l-3-6m-19 0l4 3m17-9a2 2 0 0 1 0 4m-4-4a2 2 0 0 1 0-4"/><rect width="4" height="5" x="42" y="30" fill="#f1b31c" rx="1" ry="1"/><rect width="4" height="5" x="42" y="30" fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="1" ry="1"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M52 31h4m4-7l1-1m-30 1l-1-1m13 9l-3-1m3 2l-3 1m20 14l1 1m-30-1l-1 1"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}