package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmazonShopperPanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.28 29.7c1.113-.45 3.092-1.048 3.688-.326c.644.781-.17 2.477-.92 3.794"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.798 30.223c1.759 1.397 6.954 3.535 12.488 3.535a17.003 17.003 0 0 0 10.167-3.08m-17.249-8.756a2.977 2.977 0 0 0 2.177.592h.594a1.752 1.752 0 0 0 1.75-1.753h0a1.752 1.752 0 0 0-1.75-1.754h-1.188a1.752 1.752 0 0 1-1.75-1.754h0a1.752 1.752 0 0 1 1.75-1.753h.594a2.977 2.977 0 0 1 2.177.592m4.111 3.775a2.647 2.647 0 0 0 2.647 2.647h0a2.647 2.647 0 0 0 2.647-2.647v-1.72a2.647 2.647 0 0 0-2.647-2.647h0a2.647 2.647 0 0 0-2.647 2.647m0-2.647v10.587"/>`),
		g.Group(children),
	)
}