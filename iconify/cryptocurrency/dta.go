package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m24.963 14.512l-.02-1.949L13.791 6l-2.713 1.674l-.006-.003L8 9.484v13.022l2.811 1.801h.001l.019.013v-.002L13.823 26L25 19.479l-.017-4.968zm-3.275 3.112l-7.854 4.618l-2.292-1.363l7.807-4.594zm-2.866-1.64l-4.837 2.923l.021-5.681zm-4.814-3.366l.01-2.689l7.732 4.515v2.705zM10.82 23.57l-2.334-1.395V9.896l2.346 1.401v9.572l.018-.01zm.416-3.001v-9.273l2.345-1.401v9.256zM16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-1.884-9.399l10.42-6.12l-.03 2.82l-10.437 6.136zm7.793-8.584l-10.432-6.1l2.406-1.39l10.455 6.104z"/>`),
		g.Group(children),
	)
}