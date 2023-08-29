package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fridgefriend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.656 25.395l-1.616 8.513a6.168 6.168 0 0 0 .074 2.636l1.38 5.559a1.126 1.126 0 0 1-1.092 1.397H11.598a1.126 1.126 0 0 1-1.092-1.397l1.38-5.559a6.168 6.168 0 0 0 .073-2.636l-1.615-8.513a.825.825 0 0 1 .81-.979h25.69a.825.825 0 0 1 .812.98Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.261 4.5h0a4.105 4.105 0 0 1 4.105 4.105v15.811h0h-8.21h0V8.606A4.105 4.105 0 0 1 16.26 4.5Zm-1.886 11.428l3.772-2.94m-3.772 9.063l3.772-2.941m6.547-6.566h6.546v2.496h-6.546z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.452 15.04c0 6.99-3.126 5.992-3.126 9.376m8.128-9.376c0 6.99 3.125 5.992 3.125 9.376M12.063 35.253h2.053m21.821 0h-1.96"/>`),
		g.Group(children),
	)
}