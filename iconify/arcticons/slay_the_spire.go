package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlayTheSpire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.204 43.5c-14.811-1.547-20.02-12.886-20.8-15.366a6.114 6.114 0 0 0 2.091.39s-2.905-5.705-3.65-8.858a19.433 19.433 0 0 0 2.126.283S8.606 13.465 9.172 7.122c6.024 3.756 12.402 7.724 20.48 8.717C33.693 13.287 35.819 4.5 35.819 4.5a20.298 20.298 0 0 1 3.012 10.382c.283 6.661-.166 18.2-6.626 28.618"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.71 13.289c-1.12 4.676-5.75 10.487-5.75 10.487c-7.323-.237-19.633-7.634-23.54-11.683m23.54 11.683c.472 5.48-.756 19.724-.756 19.724m4.936-18.974c-.306 1.966-2.054 4.173-2.666 6.664m-18.856-6.445c4.414 6.336 11.318 3.113 12.739 10.006c-3.671-2.927-10.663-.502-12.739-10.006Z"/>`),
		g.Group(children),
	)
}