package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.354 5.903a1 1 0 0 1 1.087-.906l5.185.472a1 1 0 1 1-.181 1.991l-5.186-.47a1 1 0 0 1-.905-1.087Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.097 12.646a1 1 0 0 1-1.086-.905l-.471-5.186a1 1 0 1 1 1.991-.181l.472 5.185a1 1 0 0 1-.906 1.087Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.828 7.172a1 1 0 0 1 0 1.414l-5.656 5.657a1 1 0 0 1-1.415-1.415l5.657-5.656a1 1 0 0 1 1.414 0Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}