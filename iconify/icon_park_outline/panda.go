package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><ellipse cx="24" cy="27" rx="18" ry="17"/><ellipse cx="16.933" cy="24.64" rx="3" ry="4" transform="rotate(15 16.933 24.64)"/><ellipse rx="3" ry="4" transform="scale(-1 1) rotate(15 -109.047 -105.16)"/><path stroke-linecap="round" d="M41 20.999c3.124-3.124 3.438-8.219.313-11.343C38.19 6.532 32.124 6.876 29 10M7 21c-3.124-3.125-3.468-8.22-.344-11.343C9.781 6.532 15.876 6.875 19 10"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 35c.5 1.294 2.2 3.277 5 .862c2.8 2.415 4.5.431 5-.862"/></g>`),
		g.Group(children),
	)
}