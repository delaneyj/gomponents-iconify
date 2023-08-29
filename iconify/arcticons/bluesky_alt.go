package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueskyAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.74 16.532v1c0 10.07-7.64 21.61-21.62 21.61a21.14 21.14 0 0 1-11.62-3.45c.6.081 1.205.118 1.81.11a15.25 15.25 0 0 0 9.44-3.24a7.56 7.56 0 0 1-7.1-5.29c.473.1.956.15 1.44.15c.676 0 1.349-.09 2-.27A7.57 7.57 0 0 1 7 19.702v-.1a7.42 7.42 0 0 0 3.44.94a7.54 7.54 0 0 1-2.39-10.06a21.58 21.58 0 0 0 15.68 7.94a6.38 6.38 0 0 1-.21-1.74a7.55 7.55 0 0 1 13.17-5.31a15.59 15.59 0 0 0 4.83-1.85a7.65 7.65 0 0 1-3.39 4.27a15.87 15.87 0 0 0 4.37-1.26a15.562 15.562 0 0 1-3.76 4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.35 25.575a2.653 2.653 0 0 1 3.176-2.6c1.111.214 1.999 1.175 2.11 2.302c.082.838-.183 1.664-.762 2.172c-1.072.94-4.524 3.476-4.524 3.476h5.3"/>`),
		g.Group(children),
	)
}