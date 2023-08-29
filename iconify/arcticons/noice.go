package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="10.6" cy="14.11" r=".75" fill="currentColor"/><circle cx="12.87" cy="9.57" r=".75" fill="currentColor"/><circle cx="17.32" cy="9.57" r=".75" fill="currentColor"/><circle cx="21.77" cy="9.57" r=".75" fill="currentColor"/><circle cx="19.53" cy="5.25" r=".75" fill="currentColor"/><circle cx="24" cy="5.25" r=".75" fill="currentColor"/><circle cx="28.47" cy="5.25" r=".75" fill="currentColor"/><circle cx="26.23" cy="9.57" r=".75" fill="currentColor"/><circle cx="30.68" cy="9.57" r=".75" fill="currentColor"/><circle cx="35.13" cy="9.57" r=".75" fill="currentColor"/><circle cx="15.07" cy="14.11" r=".75" fill="currentColor"/><circle cx="19.53" cy="14.11" r=".75" fill="currentColor"/><circle cx="24" cy="14.11" r=".75" fill="currentColor"/><circle cx="28.47" cy="14.11" r=".75" fill="currentColor"/><circle cx="32.93" cy="14.11" r=".75" fill="currentColor"/><circle cx="37.4" cy="14.11" r=".75" fill="currentColor"/><circle cx="37.4" cy="33.89" r=".75" fill="currentColor"/><circle cx="35.13" cy="38.43" r=".75" fill="currentColor"/><circle cx="30.68" cy="38.43" r=".75" fill="currentColor"/><circle cx="26.23" cy="38.43" r=".75" fill="currentColor"/><circle cx="28.47" cy="42.75" r=".75" fill="currentColor"/><circle cx="24" cy="42.75" r=".75" fill="currentColor"/><circle cx="19.53" cy="42.75" r=".75" fill="currentColor"/><circle cx="21.77" cy="38.43" r=".75" fill="currentColor"/><circle cx="17.32" cy="38.43" r=".75" fill="currentColor"/><circle cx="12.87" cy="38.43" r=".75" fill="currentColor"/><circle cx="32.93" cy="33.89" r=".75" fill="currentColor"/><circle cx="28.47" cy="33.89" r=".75" fill="currentColor"/><circle cx="24" cy="33.89" r=".75" fill="currentColor"/><circle cx="19.53" cy="33.89" r=".75" fill="currentColor"/><circle cx="15.07" cy="33.89" r=".75" fill="currentColor"/><circle cx="10.6" cy="33.89" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 24.55C6.8 23 8.5 22 10.21 22c3.15 0 3.54 3.49 6.49 3.49c2.21 0 3.47-5.47 5.52-5.47c2.37 0 2.66 8 5.19 8c1.81 0 3.39-5.56 5.65-5.56c1.92 0 2.76 2.86 5.95 2.86A8.24 8.24 0 0 0 43 24"/>`),
		g.Group(children),
	)
}