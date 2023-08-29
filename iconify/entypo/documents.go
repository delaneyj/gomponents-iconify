package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Documents(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m19.398 7.415l-7.444-1.996L10.651.558c-.109-.406-.545-.642-.973-.529L.602 2.461c-.428.114-.686.538-.577.944l3.23 12.051c.109.406.544.643.971.527l3.613-.967l-.492 1.838c-.109.406.149.83.577.943l8.11 2.174c.428.115.862-.121.972-.529l2.97-11.084c.108-.406-.15-.83-.578-.943zM1.633 3.631l7.83-2.096l2.898 10.818l-7.83 2.096L1.633 3.631zm14.045 14.832L8.864 16.6l.536-2.002l3.901-1.047c.428-.113.688-.537.578-.943l-1.508-5.627l5.947 1.631l-2.64 9.851z"/>`),
		g.Group(children),
	)
}