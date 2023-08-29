package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cakewallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.23 27.58H18v-7.74l6 6l6-6v7.74h4.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.23 27.58A11.34 11.34 0 0 1 32 16l7.2-7.2a21.47 21.47 0 0 0-15.29-6.3h0a21.5 21.5 0 1 0 15.21 36.78L32 32.07a11.35 11.35 0 0 1-17.7-2.17"/><circle cx="24" cy="24" r="11.35" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}