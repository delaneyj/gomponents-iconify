package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.92 210.62a8 8 0 1 1-11.84 10.76l-38.37-42.2l-61.86 66.28a8 8 0 0 1-13.69-7l14.66-73.33l-57.63-21.64a8 8 0 0 1-3-13l39.18-42l-39.29-43.11a8 8 0 1 1 11.84-10.76Zm-34.14-61.34a8 8 0 0 0 5.86 2.62h.06a8 8 0 0 0 5.85-2.55l22.3-23.89a8 8 0 0 0-3-12.95L153.18 90.9l14.66-73.33a8 8 0 0 0-13.69-7L108.27 59.7a8 8 0 0 0-.07 10.84Z"/>`),
		g.Group(children),
	)
}