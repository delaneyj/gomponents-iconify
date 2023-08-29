package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m163.2 33.1l-49.6 33.04L27.29 389.9l361.31 89l96.1-356l-66.3-18.5l-24.5 91.4l-105-28.1l-14.7 55.2c15.2 9.7 22.3 28.2 17.7 45.6c-5.9 22.2-28.7 35.3-50.8 29.4c-22.1-5.9-35.2-28.7-29.3-50.8c4.7-17.6 20.2-30.1 38.4-30.8c2.1-.1 4.3 0 6.4.3l14.3-53.8l-115-30.8l2.4-9l22.8-84.91zm34.9 7.78l-19.4 77.92l202.1 54.1l19.4-77.87zM70.32 331.6l312.08 74.7l-4.4 18.2l-312.04-74.7z"/>`),
		g.Group(children),
	)
}