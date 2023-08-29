package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inspection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M13 0c-.7 0-1.206.294-1.406.594c-.1.1-.181.306-.281.406H10c-1.1 0-2 .9-2 2H6C4.3 3 3 4.3 3 6v17c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V6c0-1.7-1.3-3-3-3h-2c0-1.1-.9-2-2-2h-1.313c-.1-.1-.18-.306-.28-.406C14.206.294 13.7 0 13 0zm-3 3h6v2h-6V3zM6 6h2.313c.3.6.987 1 1.687 1h6c.7 0 1.387-.4 1.688-1H20v17H6V6zm10.563 4.688c-.188 0-.37.068-.47.218l-4.405 4.688L10 14c-.3-.3-.8-.3-1 0l-.688.688c-.3.3-.3.8 0 1l2.782 2.624c.4.4 1.006.275 1.406-.125l5.406-5.78c.2-.2.207-.607-.093-.907l-.72-.594a.747.747 0 0 0-.53-.219z"/>`),
		g.Group(children),
	)
}