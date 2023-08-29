package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObservedHail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 18A12 12 0 1 0 16 6h-4V1L6 7l6 6V8h4A10 10 0 1 1 6 18Z"/><circle cx="13.5" cy="23.5" r="1.5" fill="currentColor"/><circle cx="10.5" cy="19.5" r="1.5" fill="currentColor"/><circle cx="16.5" cy="19.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M12 16.586L15.586 13L17 14.413L13.413 18zm6 0L21.586 13L23 14.413L19.413 18z"/>`),
		g.Group(children),
	)
}