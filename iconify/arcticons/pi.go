package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.5 14.824v1.426a2.775 2.775 0 0 1-2.812 2.737H15.311a2.775 2.775 0 0 0-2.811 2.737v1.427m15.717-4.163v16.577m-8.435-16.577v16.577"/><ellipse cx="19.782" cy="13.461" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.541" ry="1.5"/><ellipse cx="28.217" cy="13.461" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.541" ry="1.5"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}