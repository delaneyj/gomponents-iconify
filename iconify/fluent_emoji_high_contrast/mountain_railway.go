package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainRailway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M12.5 24a.5.5 0 0 0 0 1h1a.5.5 0 0 0 0-1h-1Z"/><path d="M30 30H2v-9.133a1 1 0 0 1 .237-.647l12.41-14.625a2 2 0 0 1 2.886-.172l2.445 2.27l4.045-4.903a2 2 0 0 1 3.03-.065l-.744.669a1 1 0 0 0-1.515.032l-4.082 4.95l1.754 1.628H29v-3.62l-2.69-2.99l.742-.67L30 6v24Zm-19-3a1 1 0 0 0 1 1a2 2 0 0 0 3.732 1H15a1 1 0 0 1-1-1h6a1 1 0 0 1-1 1h-.732A2 2 0 0 0 22 28h6a2 2 0 0 0 1 1.732V23H11v4Zm18-13a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1v-8Zm0-1v-2H12a1 1 0 0 0-1 1v1h18Zm-18 9h2a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-2v8Zm6-8a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-2Zm5 1v6a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1Z"/></g>`),
		g.Group(children),
	)
}