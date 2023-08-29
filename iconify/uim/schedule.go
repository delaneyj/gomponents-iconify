package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schedule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v2a1 1 0 0 1-.999 1H7zm10 0a1 1 0 0 1-1-1V3a1 1 0 0 1 2 0v2a1 1 0 0 1-.999 1H17z" opacity=".5"/><path fill="currentColor" d="M19 4h-1v1a1 1 0 0 1-2 0V4H8v1a1 1 0 0 1-2 0V4H5a3 3 0 0 0-3 3v2h20V7a3 3 0 0 0-3-3z"/><circle cx="7" cy="13" r="1" fill="currentColor"/><circle cx="7" cy="17" r="1" fill="currentColor"/><circle cx="12" cy="13" r="1" fill="currentColor"/><circle cx="12" cy="17" r="1" fill="currentColor"/><circle cx="17" cy="13" r="1" fill="currentColor"/><circle cx="17" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M2 9v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9H2zm5 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm5 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm5 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" opacity=".5"/>`),
		g.Group(children),
	)
}