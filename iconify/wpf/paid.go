package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M12.875 2a1 1 0 0 0-.594.313L5.687 9H1c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1h.094l2.718 9.313c.1.4.5.687 1 .687h16.5c.4 0 .8-.288 1-.688L25 13c.6 0 1-.4 1-1v-2c0-.6-.4-1-1-1h-4.594L13.72 2.281A1 1 0 0 0 12.875 2zM13 4.438L17.594 9H8.5L13 4.437zm4.563 7.343c.2 0 .38.069.53.219l.72.688c.3.3.3.824 0 1.124l-6 6c-.3.3-.825.3-1.126 0l-3.374-3.406c-.3-.3-.3-.793 0-1.093l.687-.72c.3-.3.794-.3 1.094 0l2.219 2.22L17 12a.798.798 0 0 1 .563-.219z"/>`),
		g.Group(children),
	)
}