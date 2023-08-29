package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M16.031 2.063c-.321.001-.676.145-1 .468L8.437 9H5c-.551 0-1 .449-1 1v6c0 .551.449 1 1 1h3.438L15 23.438c1 1 2 .488 2-.875V3.28c0-.791-.433-1.222-.969-1.219zm2.5 7.03a1 1 0 0 0-.031 1.97c.857.223 1.5.999 1.5 1.937c0 .938-.643 1.714-1.5 1.938a1 1 0 1 0 .5 1.937c1.721-.45 3-2.025 3-3.875s-1.279-3.425-3-3.875a1 1 0 0 0-.375-.031a1 1 0 0 0-.094 0z"/>`),
		g.Group(children),
	)
}