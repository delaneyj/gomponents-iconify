package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 1a8 8 0 0 1 3.875 15h-7.75A8 8 0 0 1 8 1zm4.53 12.53A6.364 6.364 0 0 0 14.406 9H13V8h1.329a6.346 6.346 0 0 0-.665-2H12V5h1.004a6.372 6.372 0 0 0-3.005-2.089V4h-1V2.671a6.506 6.506 0 0 0-2 0V4h-1V2.911A6.384 6.384 0 0 0 2.994 5h1.004v1H2.334a6.346 6.346 0 0 0-.665 2h1.329v1H1.592c0 1.711.666 3.32 1.876 4.53c.167.167.343.324.524.47h3.006l.571-8h.857l.571 8h3.006c.182-.146.357-.303.524-.47z"/>`),
		g.Group(children),
	)
}