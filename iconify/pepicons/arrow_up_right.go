package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.354 5.903a1 1 0 0 1 1.087-.906l5.185.472a1 1 0 1 1-.181 1.991l-5.186-.47a1 1 0 0 1-.905-1.087Z"/><path d="M14.097 12.646a1 1 0 0 1-1.086-.905l-.471-5.186a1 1 0 1 1 1.991-.181l.472 5.185a1 1 0 0 1-.906 1.087Z"/><path d="M12.828 7.172a1 1 0 0 1 0 1.414l-5.656 5.657a1 1 0 0 1-1.415-1.415l5.657-5.656a1 1 0 0 1 1.414 0Z"/></g>`),
		g.Group(children),
	)
}