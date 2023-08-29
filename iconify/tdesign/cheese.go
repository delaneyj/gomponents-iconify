package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.532 1.467l-3.997 12.658l-.833-.112a1.5 1.5 0 0 0-.89 2.82l.746.387l-1.885 5.968l-1.662-.395c-6.318-1.505-11.29-6.446-12.795-12.769L.812 8.327l10.97-3.465l.266 1.018a1.501 1.501 0 1 0 2.858-.903l-.366-.986l7.992-2.524Zm-5.537 3.846A3.502 3.502 0 0 1 13.5 9a3.496 3.496 0 0 1-2.969-1.646L3.19 9.674c1.349 5.48 5.657 9.774 11.138 11.137l.85-2.693a3.5 3.5 0 0 1 1.925-6.096l2.365-7.49l-2.472.781ZM9.5 11a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM6 12.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z"/>`),
		g.Group(children),
	)
}