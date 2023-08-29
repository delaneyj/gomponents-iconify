package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowseGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.41 1.586l.692.722A13.957 13.957 0 0 1 24 12c0 3.761-1.485 7.178-3.898 9.692l-.692.722l-1.443-1.385l.692-.721A11.956 11.956 0 0 0 22 12c0-3.225-1.27-6.15-3.34-8.308l-.693-.721l1.443-1.385ZM10 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16ZM0 12C0 6.477 4.477 2 10 2s10 4.477 10 10s-4.477 10-10 10S0 17.523 0 12Zm11-5.5v5.086L14.414 15L13 16.414l-4-4V6.5h2Z"/>`),
		g.Group(children),
	)
}