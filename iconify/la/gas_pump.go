package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasPump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 5C6.355 5 5 6.355 5 8v19h14v-7h2v4c0 1.645 1.355 3 3 3s3-1.355 3-3V13.812c0-.796-.313-1.53-.875-2.093L21.406 7L20 8.406l2.844 2.844C21.77 11.707 21 12.766 21 14c0 1.645 1.355 3 3 3a2.95 2.95 0 0 0 1-.188V24c0 .566-.434 1-1 1c-.566 0-1-.434-1-1v-4c0-1.094-.906-2-2-2h-2V8c0-1.645-1.355-3-3-3zm0 2h8c.566 0 1 .434 1 1v4H7V8c0-.566.434-1 1-1zm16 6c.563 0 1 .438 1 1c0 .563-.438 1-1 1c-.563 0-1-.438-1-1c0-.563.438-1 1-1zM7 14h10v11H7z"/>`),
		g.Group(children),
	)
}