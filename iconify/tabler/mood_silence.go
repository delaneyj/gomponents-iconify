package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodSilence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18zM9 10h-.01M15 10h-.01M8 15h8m-7-1v2m3-2v2m3-2v2"/>`),
		g.Group(children),
	)
}