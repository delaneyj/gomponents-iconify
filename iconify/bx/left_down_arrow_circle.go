package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftDownArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.006 2.007A9.93 9.93 0 0 0 4.935 4.93c-3.898 3.898-3.898 10.242 0 14.142c1.885 1.885 4.396 2.923 7.071 2.923s5.187-1.038 7.071-2.923c3.898-3.899 3.898-10.243 0-14.142a9.931 9.931 0 0 0-7.071-2.923zm5.657 15.65c-1.507 1.507-3.516 2.337-5.657 2.337s-4.15-.83-5.657-2.337c-3.118-3.119-3.118-8.194 0-11.313c1.507-1.507 3.517-2.337 5.657-2.337s4.15.83 5.657 2.337c3.118 3.119 3.118 8.194 0 11.313z"/><path fill="currentColor" d="m14.346 8.247l-3.215 3.215l-2.125-2.125V15h5.663l-2.124-2.124l3.215-3.215z"/>`),
		g.Group(children),
	)
}