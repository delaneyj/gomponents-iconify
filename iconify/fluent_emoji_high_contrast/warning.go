package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.999 22.898c-.78 0-1.4-.63-1.4-1.4v-9.17c0-.78.63-1.4 1.4-1.4c.78 0 1.4.63 1.4 1.4v9.16a1.4 1.4 0 0 1-1.4 1.41Zm1.4 2.579a1.4 1.4 0 1 1-2.8 0a1.4 1.4 0 0 1 2.8 0Z"/><path d="M13.976 5.163c.906-1.55 3.14-1.55 4.046 0l.003.005l12.664 21.937c.882 1.556-.222 3.503-2.03 3.503H3.339c-1.807 0-2.915-1.959-2.028-3.508l.002-.003l12.66-21.93l.002-.004Zm1.727 1.007v.002L3.046 28.095h-.001a.343.343 0 0 0 .293.513h25.32a.34.34 0 0 0 .291-.515L16.296 6.172l-.001-.002a.342.342 0 0 0-.592 0Z"/></g>`),
		g.Group(children),
	)
}