package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SketchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.998 2.25h.009a.65.65 0 0 1 .077.005l5.904.624a.748.748 0 0 1 .529.299l4.575 6.142a.749.749 0 0 1-.03.948L12.574 22.48a.75.75 0 0 1-1.154 0L.936 10.27a.753.753 0 0 1-.034-.947L5.477 3.18a.744.744 0 0 1 .532-.302l5.903-.624a.751.751 0 0 1 .077-.005h.009Zm9.004 6.775l-2.548-3.421l-.374 3.421h2.922Zm-7.335-5.094L16.7 7.847l.388-3.555l-3.42-.361Zm-6.759.361l3.42-.361l-3.032 3.916l-.388-3.555Zm5.09-.067l3.716 4.8H8.281l3.717-4.8Zm-8.865 6.3h3.125L9.379 17.8l-6.246-7.274Zm2.41-4.921l-2.55 3.421h2.923l-.374-3.421Zm15.32 4.921L14.617 17.8l3.121-7.274h3.125Zm-12.973 0h8.216l-4.108 9.573l-4.108-9.573Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}