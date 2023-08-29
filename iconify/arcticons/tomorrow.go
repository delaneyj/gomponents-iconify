package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tomorrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.892 4.5v33c0 2.115-2.067 6-2.067 6h12.299s-2.066-3.885-2.066-6v-33Zm8.222 0c5.68.932 13.178 9.739 13.178 9.739L39.347 4.5Zm-8.227 0c-5.681.932-13.179 9.739-13.179 9.739L8.653 4.5Z"/>`),
		g.Group(children),
	)
}