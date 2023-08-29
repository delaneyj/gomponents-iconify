package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Garlic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.91 4.17c0-1.202.969-2.17 2.17-2.17h5.81c1.203 0 2.17.968 2.17 2.17v1.5c0 2.242.998 4.358 2.726 5.777l.001.002l4.23 3.49c5.854 4.838 2.43 14.331-5.156 14.331H9.11c-7.586 0-11.01-9.493-5.157-14.33l4.23-3.491l.002-.002A7.466 7.466 0 0 0 10.91 5.67v-1.5Zm8.81 7.71a7.595 7.595 0 0 1-2.35-5.49V3h-2.81v3.42c0 2.08-.85 4.06-2.36 5.49l-2.29 2.18c-5.38 5.12-1.76 14.19 5.67 14.19h.79c7.44 0 11.06-9.06 5.68-14.19l-2.33-2.21Z"/>`),
		g.Group(children),
	)
}