package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 601V169s54-81 196-81c143 0 188 77 188 77s45-77 188-77c142 0 196 81 196 81v432s-95-70-198-70c-102 0-186 75-186 75s-84-75-186-75C95 531 0 601 0 601zm65-415v314s41-31 139-31c97 0 148 58 148 58V200s-69-50-156-50c-86 0-131 36-131 36zm351 14v327s48-58 148-58c99 0 139 31 139 31V186s-45-36-131-36c-87 0-156 50-156 50z"/>`),
		g.Group(children),
	)
}