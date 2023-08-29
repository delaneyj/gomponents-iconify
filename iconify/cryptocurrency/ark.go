package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-.053-18.653L27 25L15.996 7L5 24.891l10.947-11.544zm1.588 4.585v.017l-1.662-1.953l-1.76 1.936h3.422zm-6.6 3.177h9.859l-1.998-2.045l-5.92.025v.009l-1.94 1.987v.024z"/>`),
		g.Group(children),
	)
}