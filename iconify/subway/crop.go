package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 21.3V0h-21.3l-85.3 85.3H128V0H85.3v85.3H0V128h85.3v298.7H384V512h42.7v-85.3H512V384h-85.3V106.7L512 21.3zM128 128h234.7L128 362.7V128zm256 256H149.3L384 149.3V384z"/>`),
		g.Group(children),
	)
}