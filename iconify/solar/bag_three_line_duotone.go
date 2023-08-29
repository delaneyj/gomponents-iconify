package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagThreeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.742 18.555C4.942 20 7.174 20 11.639 20h.722c4.465 0 6.698 0 7.897-1.445m-16.516 0c-1.2-1.446-.789-3.64.034-8.03c.586-3.12.878-4.681 1.99-5.603M3.741 18.555Zm16.516 0c1.2-1.446.788-3.64-.034-8.03c-.586-3.12-.878-4.681-1.99-5.603m2.024 13.633ZM18.235 4.922C17.124 4 15.536 4 12.36 4h-.722c-3.175 0-4.763 0-5.874.922m12.47 0Zm-12.47 0Z"/><path stroke-linecap="round" d="M9.17 8a3.001 3.001 0 0 0 5.66 0" opacity=".5"/></g>`),
		g.Group(children),
	)
}