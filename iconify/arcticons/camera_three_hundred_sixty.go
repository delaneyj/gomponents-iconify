package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraThreeHundredSixty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.327" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="29.818" r="11.108" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.009 8.832c-1.925 1.397-2.177 5.939.767 7.411s5.803.42 7.318-1.598s11.143-13.113 18.293-8.68"/><circle cx="12.678" cy="11.995" r="2.439" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.622" cy="26.439" r="2.117" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="29.818" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}