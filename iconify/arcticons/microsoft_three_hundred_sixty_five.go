package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftThreeHundredSixtyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.507 3.241L8.443 9.74a8.599 8.599 0 0 0-4.508 7.563V30.7a8.599 8.599 0 0 0 4.508 7.564m15.668-5.598l-2.826-1.556a8.599 8.599 0 0 1-4.45-7.532v-4.072"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.166 19.275v4.45a8.599 8.599 0 0 1-4.508 7.564l-11.466 6.202a8.599 8.599 0 0 1-7.435.358c.22.148.45.286.687.414l11.465 6.202a8.599 8.599 0 0 0 8.182 0l11.465-6.202a8.599 8.599 0 0 0 4.508-7.564"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.557 9.739L28.092 3.536a8.6 8.6 0 0 0-7.585-.295a8.6 8.6 0 0 0-3.673 7.048v8.986l3.288-1.661a8.6 8.6 0 0 1 7.756 0l11.465 5.793a8.6 8.6 0 0 1 4.72 7.484a7.59 7.59 0 0 0 .001-.191V17.302a8.6 8.6 0 0 0-4.507-7.563h0Z"/>`),
		g.Group(children),
	)
}