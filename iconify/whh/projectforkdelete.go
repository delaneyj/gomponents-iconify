package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projectforkdelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m951.85 161l64 64q9 9 9 21.5t-9 21.5l-43 44q-9 9-21.5 9t-21.5-9l-65-65l-65 66q-9 9-21.5 9t-21.5-9l-44-44q-9-9-9-21.5t9-21.5l66-65l-66-66q-9-8-9-21t9-22l44-43q9-9 21.5-9t21.5 9l65 65l65-64q9-9 21.5-9t21.5 9l43 43q9 9 9 21.5t-9 21.5zm-752 631q57 39 57 105q0 53-37.5 90.5t-90.5 37.5t-90.5-37.5T.85 897q0-35 17.5-64t46.5-46V65q0-27 18.5-45.5T128.85 1t45.5 18.5t18.5 45.5v370q78-122 194-201.5t254-104.5v128q-174 39-294.5 183t-146.5 352zm-70.5 41q-26.5 0-45.5 18.5t-19 45t18.5 45.5t45.5 19t45.5-19t18.5-45.5t-18.5-45t-45-18.5z"/>`),
		g.Group(children),
	)
}