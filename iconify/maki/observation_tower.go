package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObservationTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12 12.5h-.5L10 9V7c.995 0 .995-.75.995-.75L11.5 3s0-.5-.5-.5h-1s0-.5-.5-.5h-1v-.5s0-1-1-1s-1 1-1 1V2h-1c-.5 0-.5.5-.5.5H4c-.5 0-.5.5-.5.5l.505 3.25S4.005 7 5 7v2l-1.5 3.5H3s-1 0-1 .75S3 14 3 14h9s1 0 1-.75s-1-.75-1-.75Zm-1.75-9L10 5H5l-.25-1.503l5.5.003ZM8.5 7v1h-2V7h2Zm-2 2.497h2l1 3.003h-4l1-3.003Z"/>`),
		g.Group(children),
	)
}