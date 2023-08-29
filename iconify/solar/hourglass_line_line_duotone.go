package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassLineLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12 12l-2.958 2.929c-2.922 2.894-4.383 4.341-3.974 5.59c.035.107.078.211.13.312C5.8 22 7.867 22 12 22c4.133 0 6.2 0 6.802-1.17c.052-.1.095-.204.13-.311c.41-1.249-1.052-2.696-3.974-5.59L12 12Zm0 0l2.958-2.929c2.922-2.894 4.383-4.341 3.974-5.59a2.12 2.12 0 0 0-.13-.312C18.2 2 16.133 2 12 2C7.867 2 5.8 2 5.198 3.17c-.052.1-.095.204-.13.311c-.41 1.249 1.052 2.696 3.974 5.59L12 12Z" opacity=".5"/><path stroke-linecap="round" d="M10 5.5h4m-4 13h4"/></g>`),
		g.Group(children),
	)
}