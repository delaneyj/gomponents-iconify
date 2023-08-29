package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Codrops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#237DAC" d="M256 128c0 70.692-57.308 128-128 128C57.307 256 0 198.692 0 128C0 57.307 57.307 0 128 0c70.692 0 128 57.307 128 128"/><path fill="#FFF" d="M129.283 60.731s39.839 64.87 39.839 88.5c0 18.839-8.81 41.505-39.839 41.505c-31.029 0-39.854-22.679-39.854-41.517c0-22.818 39.854-88.488 39.854-88.488"/><path fill="#61AAD4" d="M100.532 82.518s25.074 40.828 25.074 55.701c0 11.857-5.545 26.123-25.074 26.123c-19.53 0-25.084-14.274-25.084-26.131c0-14.361 25.084-55.693 25.084-55.693"/>`),
		g.Group(children),
	)
}