package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hammer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m15.781 12.953l-4.712-4.712a.752.752 0 0 0-1.061 0l-.354.354L6.779 5.72L11.499 1h-5l-2.22 2.22l-.22-.22H2.998v1.061l.22.22l-3.22 3.22l2.5 2.5l3.22-3.22l2.875 2.875l-.354.354a.752.752 0 0 0 0 1.061l4.712 4.712a.752.752 0 0 0 1.061 0l1.768-1.768a.752.752 0 0 0 0-1.061z"/>`),
		g.Group(children),
	)
}