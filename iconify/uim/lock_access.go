package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10c-.6 0-1-.4-1-1V4h-5c-.6 0-1-.4-1-1s.4-1 1-1h6c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1zM3 10c-.6 0-1-.4-1-1V3c0-.6.4-1 1-1h6c.6 0 1 .4 1 1s-.4 1-1 1H4v5c0 .6-.4 1-1 1zm6 12H3c-.6 0-1-.4-1-1v-6c0-.6.4-1 1-1s1 .4 1 1v5h5c.6 0 1 .4 1 1s-.4 1-1 1zm12 0h-6c-.6 0-1-.4-1-1s.4-1 1-1h5v-5c0-.6.4-1 1-1s1 .4 1 1v6c0 .6-.4 1-1 1z"/><path fill="currentColor" d="M9 10h6c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2H9c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2zm2-1c0-.6.4-1 1-1s1 .4 1 1v1h2V9c0-1.7-1.3-3-3-3S9 7.3 9 9v1h2V9z" opacity=".5"/>`),
		g.Group(children),
	)
}