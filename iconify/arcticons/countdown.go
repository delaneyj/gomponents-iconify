package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Countdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5c11.874 0 21.5 9.626 21.5 21.5S35.874 45.5 24 45.5S2.5 35.874 2.5 24m18 7h7m-7-12.094L24 17m0 0v14"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M2.5 24c0-.504.017-1.004.051-1.5"/><path stroke-dasharray="0 0 2.797 2.797" d="M2.925 19.728c1.794-8.902 9.087-15.812 18.183-17.035"/><path d="M22.5 2.551c.496-.034.996-.051 1.5-.051"/></g>`),
		g.Group(children),
	)
}