package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 6a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-5 4a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 8a.5.5 0 0 1 .5.5V10a.5.5 0 0 1-1 0V8.5A.5.5 0 0 1 10 8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12 10a.5.5 0 0 1-.5.5H10a.5.5 0 0 1 0-1h1.5a.5.5 0 0 1 .5.5ZM8 2.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.637 2.02a.5.5 0 0 1 .344.617l-1 3.5a.5.5 0 0 1-.962-.274l1-3.5a.5.5 0 0 1 .618-.344Zm0 15.96a.5.5 0 0 0 .344-.617l-1-3.5a.5.5 0 0 0-.962.274l1 3.5a.5.5 0 0 0 .618.344Zm2.726 0a.5.5 0 0 1-.344-.617l1-3.5a.5.5 0 0 1 .962.274l-1 3.5a.5.5 0 0 1-.618.344Zm0-15.96a.5.5 0 0 0-.344.617l1 3.5a.5.5 0 0 0 .962-.274l-1-3.5a.5.5 0 0 0-.618-.344Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 17.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}