package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerScreenTvMoviesTelevisionCathodeCrtTvRayTubeVintageVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="10.5" x=".5" y="3" rx="1"/><rect width="8" height="5.5" x="3" y="5.5" rx="1"/><path d="M5 .5L7 3L9 .5"/></g>`),
		g.Group(children),
	)
}