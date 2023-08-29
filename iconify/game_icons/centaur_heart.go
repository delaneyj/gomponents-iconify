package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CentaurHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M147.9 25.75c-10.1 0-20.2 1.2-29.9 3.6c-52 13-95.7 59.8-97.8 127.05c-1.1 35.8 9.9 65.6 27.6 91.1c17.6 25.6 41.5 47.1 66.8 68.5C165.2 358.5 221 399.8 247 464.3l8.9 22l8.5-22.2C289 399.5 343 356.3 393 312.8c50-43.5 96.6-88.3 98.8-155.8c2.3-71.55-42.4-116.75-95.5-127.15c-49-9.7-105.4 9.1-140.3 57.7c-27.5-42.1-68.3-61.7-108.1-61.8zM80 112l80 32l-64 32l116.4 23.3L256 112l43.6 87.3L416 176l-64-32l80-32l32 80l-171.5 66L256 416l-36.5-158L48 192l32-80z"/>`),
		g.Group(children),
	)
}