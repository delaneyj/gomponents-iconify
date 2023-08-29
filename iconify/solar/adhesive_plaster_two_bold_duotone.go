package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdhesivePlasterTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m20.416 12.765l-9.181-9.18a5.41 5.41 0 0 0-7.65 7.65l9.18 9.18a5.41 5.41 0 1 0 7.65-7.65Z" opacity=".5"/><path d="m19.885 12.234l.531.531c.18.18.344.37.494.566l-7.578 7.578a5.45 5.45 0 0 1-.567-.493l-.53-.531l7.65-7.65ZM10.668 3.09c.198.15.387.314.567.494l.53.53l-7.65 7.651l-.53-.53c-.18-.18-.345-.37-.494-.567l7.577-7.578Z"/><circle cx="9.172" cy="12" r="1" transform="rotate(-45 9.172 12)"/><circle cx="12" cy="14.829" r="1" transform="rotate(-45 12 14.829)"/><circle cx="12" cy="9.171" r="1" transform="rotate(-45 12 9.171)"/><circle cx="14.828" cy="12" r="1" transform="rotate(-45 14.828 12)"/></g>`),
		g.Group(children),
	)
}