package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kyivdigital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.013 33.553L24.027 43.5l-17.04-9.947V4.5h34.026v29.053zM17 10.92v18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31 28.92l-12.724-9L31 10.98m-12.724 8.94H17"/>`),
		g.Group(children),
	)
}