package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Android(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.06 9.33a.94.94 0 0 0-.94.94v4.26a.94.94 0 0 0 1.88 0v-4.26a.94.94 0 0 0-.94-.94Zm-12.12 0a.94.94 0 0 0-.94.94v4.26a.94.94 0 0 0 1.88 0v-4.26a.94.94 0 0 0-.94-.94Zm1.62 0v6.4a1.07 1.07 0 0 0 1.07 1.07h.68v2.27a.94.94 0 0 0 1.88 0V16.8h1.62v2.27a.94.94 0 0 0 1.88 0V16.8h.68a1.07 1.07 0 0 0 1.07-1.07v-6.4Z"/><circle cx="9.98" cy="7.07" r=".4" fill="none"/><circle cx="14.02" cy="7.07" r=".4" fill="none"/><path fill="currentColor" d="M16.32 8a4.13 4.13 0 0 0-1.83-2.36l-.15-.09l-.16-.08l.18-.31l.53-1a.14.14 0 0 0-.05-.16h-.07a.17.17 0 0 0-.13.07L14.1 5l-.17.31l-.16-.07l-.17-.06a4.88 4.88 0 0 0-3.2 0l-.16.06l-.17.07L9.9 5l-.54-1a.14.14 0 0 0-.25.14l.53 1l.18.31l-.16.08l-.15.09A4.07 4.07 0 0 0 7.69 8a3.11 3.11 0 0 0-.13.8h8.88a3.6 3.6 0 0 0-.12-.8ZM10 7.47a.4.4 0 1 1 .4-.4a.4.4 0 0 1-.4.4Zm4 0a.4.4 0 1 1 .4-.4a.4.4 0 0 1-.4.4Z"/>`),
		g.Group(children),
	)
}