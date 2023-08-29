package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.781 8.132A3.424 3.424 0 1 0 15.94 3.29l-1.748 1.748l-1.286-1.285l-1.414 1.414l1.286 1.285L2.072 17.157v4.842h4.843l10.704-10.704l1.285 1.286l1.415-1.415l-1.286-1.285l1.748-1.749Zm-4.576 1.749L6.086 19.999H4.072v-2.014L14.191 7.867l2.014 2.014Zm-.6-3.429l1.748-1.748a1.424 1.424 0 1 1 2.014 2.014L17.62 8.466l-2.014-2.014Z"/>`),
		g.Group(children),
	)
}