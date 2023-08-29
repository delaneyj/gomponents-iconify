package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOrange0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M41 26c0 9.941-6 18-17 18S7 35.94 7 26c0-3.68 1.104-7.102 3-9.953C13.225 11.197 17.74 13 24 13s10.775-1.803 14 3.047c1.895 2.85 3 6.273 3 9.953Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26 13l3-4h-3l-2-2l-2 2h-3l3 4"/><circle cx="18" cy="20" r="2" fill="#000"/><circle cx="15" cy="27" r="2" fill="#000"/><circle cx="21" cy="25" r="2" fill="#000"/><circle cx="18" cy="32" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOrange0)"/>`),
		g.Group(children),
	)
}