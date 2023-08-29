package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rattle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRattle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="30.075" cy="15.562" r="11" fill="#fff" transform="rotate(40 30.075 15.562)"/><path d="M21.648 8.491c-4.61.049-6.326-3.35-8.576-.67c-2.032 2.423.493 5.636-1.435 7.934m13.665 10.162l-5.785 6.894c-1.607 1.915-1.682 5.116-3.61 7.414s-4.746 2.545-7.044.616c-2.298-1.928-2.545-4.746-.617-7.044c1.929-2.298 5.068-2.928 6.675-4.843c1.607-1.915 5.163-5.893 5.785-6.894"/><circle cx="11.24" cy="19.339" r="3" fill="#fff" transform="rotate(40 11.24 19.34)"/><circle cx="28.462" cy="37.707" r="3" fill="#fff" transform="rotate(40 28.462 37.707)"/><path d="M37.216 24.165c.63 3.139 2.853 8.268.863 9.862c-1.99 1.593-6.525-1.559-7.687 1.382"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRattle0)"/>`),
		g.Group(children),
	)
}