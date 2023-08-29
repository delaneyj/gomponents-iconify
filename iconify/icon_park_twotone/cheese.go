package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCheese0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M43 20c0-2.172-18.108-11.888-22.866-14.404a1.974 1.974 0 0 0-2.149.201L5 16"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 17.652c0-1.329 1.269-2.298 2.555-1.964c6.407 1.662 23.305 5.645 32.606 4.07c1.395-.235 2.839.779 2.839 2.193v16.151a2 2 0 0 1-1.895 1.998l-34 1.79A2 2 0 0 1 5 39.891v-22.24Z"/><circle cx="12" cy="25" r="2" fill="#fff"/><circle cx="25" cy="27" r="2" fill="#fff"/><circle cx="34" cy="32" r="2" fill="#fff"/><circle cx="18" cy="32" r="2" fill="#fff" stroke="#fff" stroke-width="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCheese0)"/>`),
		g.Group(children),
	)
}