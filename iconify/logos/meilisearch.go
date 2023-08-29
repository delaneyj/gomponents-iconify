package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meilisearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosMeilisearch0" x1="153.821%" x2="19.172%" y1="-7.638%" y2="89.239%"><stop offset="0%" stop-color="#FF5CAA"/><stop offset="100%" stop-color="#FF4E62"/></linearGradient><linearGradient id="logosMeilisearch1" x1="117.325%" x2="-17.323%" y1="-7.638%" y2="89.238%"><stop offset="0%" stop-color="#FF5CAA"/><stop offset="100%" stop-color="#FF4E62"/></linearGradient><linearGradient id="logosMeilisearch2" x1="80.828%" x2="-53.821%" y1="-7.638%" y2="89.238%"><stop offset="0%" stop-color="#FF5CAA"/><stop offset="100%" stop-color="#FF4E62"/></linearGradient></defs><path fill="url(#logosMeilisearch0)" d="M0 149.288L47.297 28.277A44.462 44.462 0 0 1 88.708 0h28.515L69.926 121.012a44.462 44.462 0 0 1-41.411 28.276H0Z"/><path fill="url(#logosMeilisearch1)" d="m69.386 149.289l47.297-121.012A44.462 44.462 0 0 1 158.095 0h28.514l-47.297 121.012a44.462 44.462 0 0 1-41.411 28.277H69.386Z"/><path fill="url(#logosMeilisearch2)" d="m138.777 149.289l47.297-121.012A44.46 44.46 0 0 1 227.484 0H256l-47.297 121.012a44.463 44.463 0 0 1-41.412 28.277h-28.514Z"/>`),
		g.Group(children),
	)
}