package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PetroleumSolventSteam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M245 365h-44V151h121v37h-77v55h73v36h-73v86zm11 147q106 0 181-75t75-181t-75-181T256 0T75 75T0 256t75 181t181 75zm0-469q88 0 150.5 62.5T469 256t-62.5 150.5T256 469t-150.5-62.5T43 256t62.5-150.5T256 43z"/>`),
		g.Group(children),
	)
}