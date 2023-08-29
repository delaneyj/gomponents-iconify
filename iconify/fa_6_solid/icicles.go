package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icicles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M75.8 304.8L1 35.7c-.7-2.5-1-5-1-7.5C0 12.6 12.6 0 28.2 0h454.2C498.8 0 512 13.2 512 29.6c0 1.6-.1 3.3-.4 4.9l-77 461.6c-1.5 9.2-9.5 15.9-18.8 15.9c-9.2 0-17.1-6.6-18.7-15.6L336 160l-28.8 143.9c-1.9 9.3-10.1 16.1-19.6 16.1c-9.2 0-17.2-6.2-19.4-15.1L240 192l-29.4 176.2c-1.5 9.1-9.4 15.8-18.6 15.8s-17.1-6.7-18.6-15.8L144 192l-28.1 112.3c-2.3 9.2-10.6 15.7-20.1 15.7c-9.3 0-17.5-6.2-20-15.2z"/>`),
		g.Group(children),
	)
}