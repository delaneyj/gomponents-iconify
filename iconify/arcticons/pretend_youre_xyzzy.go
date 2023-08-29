package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PretendYoureXyzzy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.546 13.97L19.17 6.553a1.606 1.606 0 0 0-2.058.96h0L7.497 33.928a1.606 1.606 0 0 0 .96 2.058s0 0 0 0l20.377 7.417a1.606 1.606 0 0 0 2.058-.96s0 0 0 0l9.614-26.415a1.606 1.606 0 0 0-.96-2.058s0 0 0 0Zm-3.194-1.163l.469-2.656a1.606 1.606 0 0 0-1.303-1.86h0L14.163 4.524a1.606 1.606 0 0 0-1.861 1.303h0L7.421 33.51c-.04.227-.03.46.027.682m4.737-27.71H9.003c-.887 0-1.607.72-1.607 1.607v28.109c0 .887.72 1.606 1.607 1.606h4.448M32.248 7.712a1.606 1.606 0 0 0-1.561-1.23h-5.424"/>`),
		g.Group(children),
	)
}