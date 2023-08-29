package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 19.796l3-5.197h12l-3 5.197H6z"/><path fill="currentColor" d="M15 14.599h6L15 4.204H9L15 14.6z" opacity=".25"/><path fill="currentColor" d="m3 14.599l3 5.197L12 9.4L9 4.204L3 14.6z" opacity=".5"/>`),
		g.Group(children),
	)
}