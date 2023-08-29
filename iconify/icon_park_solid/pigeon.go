package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigeon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPigeon0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 26c-1.04-1.793-2.15-5.551 2.008-10.244c1.213-1.141 2.806-2.64 5.716-3.423C24.842 10.867 26.401 8.52 24.323 5C25.882 5.978 29 9.693 29 16.733"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16.253 27.93C8 23.57 4.51 30.195 4 33.755c4 0 8.679 2.911 10.721 5.823c3.676 4.66 9.36 3.56 11.742 2.427c7.352-3.883 9.87-3.56 10.21-2.912c.41 3.106 1.532 3.883 2.043 3.883c3.676.388 4.935-4.045 5.105-6.31c.817-9.319-1.361-9.707-2.552-8.736c-4.902 5.824-7.829 6.957-8.68 6.795c-3.675-.777-3.233-2.265-2.552-2.913C36.572 26.377 36.504 14.34 35.653 9c-2.45 5.825-6.467 8.251-8.169 8.737c-10.21 2.718-11.742 7.928-11.231 10.193Z"/><circle cx="12" cy="31.412" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPigeon0)"/>`),
		g.Group(children),
	)
}