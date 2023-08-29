package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shopsy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17.915 35.5c1.35 1.135 2.81 1.655 6.085 1.655h1.66a4.897 4.897 0 0 0 4.891-4.902h0a4.897 4.897 0 0 0-4.891-4.902h-3.32a4.897 4.897 0 0 1-4.891-4.902h0a4.897 4.897 0 0 1 4.891-4.902H24c3.276 0 4.734.52 6.085 1.654m-14.099-8.017h15.577"/><path d="M23.962 43.488H11.058c-4.411.266-4.082-4.148-4.082-4.148l3.487-24.67c.266-3.614 3.344-3.482 3.344-3.482l3.718-.003s-1.267-6.421 6.27-6.684h.482h-.554h.483c7.536.263 6.269 6.683 6.269 6.683l3.718.005s3.079-.133 3.344 3.48l3.487 24.67s.33 4.415-4.082 4.15h-12.98Z"/></g>`),
		g.Group(children),
	)
}