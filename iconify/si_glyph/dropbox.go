package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.98 1.41L2.615 4.36L.042 2.695L4.627.015L7.98 1.41zm7.999 1.285L12 0L9.016 1.388l4.944 3.139l2.019-1.832zM6.13 12.667L2 10.184v2.455L7.953 16v-5.675L6.13 12.667z"/><path d="M2.589 5.333L1 7.723l5.115 2.96l1.838-2.451l-5.364-2.899zM9 10.065V16l5.979-3.167v-2.472l-4.134 2.047L9 10.065z"/><path d="m9.021 8.22l1.864 2.45l5.094-2.726l-1.515-2.624l-5.443 2.9z"/></g>`),
		g.Group(children),
	)
}