package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneBlurOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="14" cy="10" r="1.5" fill="currentColor"/><circle cx="14" cy="18" r="1" fill="currentColor"/><circle cx="14" cy="14" r="1.5" fill="currentColor"/><circle cx="14" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M3 9.5c-.28 0-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5zM14.5 3c0-.28-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5zM21 14.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5z"/><circle cx="18" cy="18" r="1" fill="currentColor"/><path fill="currentColor" d="M13.5 21c0 .28.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5zM21 10.5c.28 0 .5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5z"/><circle cx="18" cy="14" r="1" fill="currentColor"/><circle cx="18" cy="6" r="1" fill="currentColor"/><circle cx="6" cy="18" r="1" fill="currentColor"/><circle cx="6" cy="14" r="1" fill="currentColor"/><path fill="currentColor" d="M3.5 14c0-.28-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5z"/><circle cx="10" cy="6" r="1" fill="currentColor"/><circle cx="6" cy="10" r="1" fill="currentColor"/><circle cx="6" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M9.5 21c0 .28.22.5.5.5s.5-.22.5-.5s-.22-.5-.5-.5s-.5.22-.5.5z"/><circle cx="10" cy="18" r="1" fill="currentColor"/><path fill="currentColor" d="M10.5 3c0-.28-.22-.5-.5-.5s-.5.22-.5.5s.22.5.5.5s.5-.22.5-.5z"/><circle cx="10" cy="14" r="1.5" fill="currentColor"/><circle cx="10" cy="10" r="1.5" fill="currentColor"/><circle cx="18" cy="10" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}