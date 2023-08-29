package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cattle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCattle0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-width="4" d="M11.465 19.923C10.682 12.481 16.517 6 24 6c7.482 0 13.317 6.481 12.534 13.923l-1.488 14.132a11.108 11.108 0 0 1-22.093 0l-1.488-14.132Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 31c3.492-3.125 11.428-7.17 20 0"/><circle cx="19" cy="18" r="2" fill="#000"/><circle cx="21" cy="34" r="2" fill="#000"/><circle cx="29" cy="18" r="2" fill="#000"/><circle cx="27" cy="34" r="2" fill="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 7.913c2.167-2 7.392-5.608 10-3c2.608 2.607 0 5-2 5.5c-2.5.625-4.2 2.3-5 3.5m-20.904-6c-2.166-2-7.392-5.608-10-3c-2.608 2.607 0 5 2 5.5c2.5.625 4.2 2.3 5 3.5"/><path stroke="#fff" stroke-width="4" d="m12 25l1 9.5M36 25l-1 9.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCattle0)"/>`),
		g.Group(children),
	)
}