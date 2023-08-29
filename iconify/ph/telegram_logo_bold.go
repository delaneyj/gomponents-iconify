package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelegramLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M239.49 23.16a13 13 0 0 0-13.23-2.26L23.6 100.21a18.22 18.22 0 0 0 3.12 34.86L76 144.74V200a20 20 0 0 0 34.4 13.88l22.67-23.51L170.35 223a20 20 0 0 0 32.7-10.54l40.62-176.55a13 13 0 0 0-4.18-12.75Zm-92.08 54.36l-62.19 44.57l-34.43-6.75ZM100 190.06v-28.71l15 13.15Zm81.16 10.52l-73.88-64.77l106.31-76.18Z"/>`),
		g.Group(children),
	)
}