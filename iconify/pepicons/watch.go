package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10 6.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM4.5 10a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z"/><path d="M10 7.75a.75.75 0 0 1 .75.75V10a.75.75 0 0 1-1.5 0V8.5a.75.75 0 0 1 .75-.75Z"/><path d="M12.25 10a.75.75 0 0 1-.75.75H10a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 .75.75ZM7.5 2.5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/><path d="M8.775 1.538a1 1 0 0 1 .687 1.237l-1 3.5a1 1 0 1 1-1.924-.55l1-3.5a1 1 0 0 1 1.237-.687Zm0 16.924a1 1 0 0 0 .687-1.237l-1-3.5a1 1 0 1 0-1.924.55l1 3.5a1 1 0 0 0 1.237.687Zm2.45 0a1 1 0 0 1-.686-1.237l1-3.5a1 1 0 1 1 1.922.55l-1 3.5a1 1 0 0 1-1.236.687Zm0-16.924a1 1 0 0 0-.686 1.237l1 3.5a1 1 0 1 0 1.922-.55l-1-3.5a1 1 0 0 0-1.236-.687Z"/><path d="M7.5 17.5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/></g>`),
		g.Group(children),
	)
}