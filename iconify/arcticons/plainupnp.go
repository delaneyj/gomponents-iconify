package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plainupnp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.17 27.971l1.52-1.504a1.772 1.772 0 0 0 .002-2.505l-.002-.002h0a1.771 1.771 0 0 0-2.505-.012l-.013.012l-1.515 1.504a10.807 10.807 0 0 0-12.634 0l-1.46-1.481a1.783 1.783 0 0 0-2.52-.003l-.003.003h0a1.777 1.777 0 0 0 0 2.506l1.516 1.493A10.829 10.829 0 0 0 5.5 34.31v5.175c0 .631.511 1.142 1.142 1.142h19.441a1.142 1.142 0 0 0 1.114-1.142V34.31a10.813 10.813 0 0 0-2.027-6.339ZM11.8 36.043a2.005 2.005 0 1 1 2-2.011v.011a2 2 0 0 1-2 2h-.011h.011Zm9.147 0a2.005 2.005 0 1 1 1.989-2.022v.022a2 2 0 0 1-2 2h.01Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.522 15.393c10.47 0 18.957 8.488 18.957 18.958"/><path d="M15.522 7.372c14.9 0 26.978 12.08 26.978 26.979"/></g>`),
		g.Group(children),
	)
}