package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialAdobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M59.386 18.139L86 81.861V18.139zm-45.386 0v63.722l26.635-63.722zm24.373 50.904h12.409l5.075 12.814H66.97L50.01 41.622z"/>`),
		g.Group(children),
	)
}