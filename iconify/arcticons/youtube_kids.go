package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeKids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.527 35.361l-7.396.53l-.05.003a25.32 25.32 0 0 0-7.056 1.638l-4.759 1.794a7.506 7.506 0 0 1-9.982-5.422L4.674 21.95a7.506 7.506 0 0 1 5.733-8.934l21.325-4.655a7.506 7.506 0 0 1 8.934 5.732l2.659 12.18a7.506 7.506 0 0 1-6.798 9.087Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.164 18.952l2.316 10.61l9.593-6.993l-11.909-3.617z"/>`),
		g.Group(children),
	)
}