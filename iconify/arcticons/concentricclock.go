package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Concentricclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="26.1" cy="24" r=".5" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.8 22.2a3.714 3.714 0 0 1 3.3-2a4.172 4.172 0 0 1 3.4 1.8m-3.6-5.6a7.602 7.602 0 0 1 6.7 3.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.8 32.4a11.407 11.407 0 0 1-19-9.8a11.525 11.525 0 0 1 8.7-9.7a11.397 11.397 0 0 1 12.1 4.9m-2.7-7.3a15.577 15.577 0 0 1 5.9 5.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.5 28.8a18.912 18.912 0 1 1-2.4-15.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 21.6A22.844 22.844 0 0 1 9.7 8.2"/>`),
		g.Group(children),
	)
}