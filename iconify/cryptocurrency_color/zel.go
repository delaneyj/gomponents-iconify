package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#183C87" fill-rule="nonzero"/><path fill="#FFF" d="m5 15.615l4.495-2.538l4.494 2.538v4.847L9.495 23L5 20.462v-4.847zm13.01 0l4.495-2.538L27 15.615v4.847L22.505 23l-4.494-2.538v-4.847zm-.472 4.21l-1.577.867l-1.499-.823V15.23l-4.1-2.316V11.23L15.961 8l5.598 3.23v1.73l-4.021 2.27v4.596z"/></g>`),
		g.Group(children),
	)
}