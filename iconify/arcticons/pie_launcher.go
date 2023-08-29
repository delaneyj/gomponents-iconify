package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.957" cy="35.322" r="6.289" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="39.343" cy="18.63" r="5.157" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.006" cy="39.225" r="5.157" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="7.525" cy="29.494" r="4.025" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.78" cy="7.643" r="4.025" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="rotate(-.439 29.78 7.645)"/><circle cx="18.253" cy="7.045" r="2.893" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="7.016" cy="18.356" r="2.893" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.258" cy="11.435" r="1.761" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}