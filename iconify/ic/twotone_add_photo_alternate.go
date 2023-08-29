package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneAddPhotoAlternate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.21 16.83l-1.96-2.36L5.5 18h11l-3.54-4.71z"/><path fill="currentColor" d="M16.5 18h-11l2.75-3.53l1.96 2.36l2.75-3.54L16.5 18zM17 7h-3V6H4v14h14V10h-1V7z" opacity=".3"/><path fill="currentColor" d="M20 4V1h-2v3h-3v2h3v2.99h2V6h3V4zm-2 16H4V6h10V4H4c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V10h-2v10z"/>`),
		g.Group(children),
	)
}