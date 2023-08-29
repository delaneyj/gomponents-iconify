package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 66h-26V56a14 14 0 0 0-14-14H64a14 14 0 0 0-14 14v10H24a14 14 0 0 0-14 14v16a38 38 0 0 0 38 38h5.14A78 78 0 0 0 122 189.75V218H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12h-26v-28.25c32.44-2.52 59.43-25.3 68.62-55.75H208a38 38 0 0 0 38-38V80a14 14 0 0 0-14-14ZM48 122a26 26 0 0 1-26-26V80a2 2 0 0 1 2-2h26v34a80.87 80.87 0 0 0 .65 10Zm80.5 56A66 66 0 0 1 62 112V56a2 2 0 0 1 2-2h128a2 2 0 0 1 2 2v55.1c0 36.62-29.38 66.63-65.5 66.9ZM234 96a26 26 0 0 1-26 26h-2.77a78.45 78.45 0 0 0 .77-10.9V78h26a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}