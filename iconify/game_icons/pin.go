package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M326.953 22.87L306.68 83.685l20.273 20.273l-115.428 115.427c-16.39-8-34.277-14.452-51.84-18.502c-14.247-3.285-28.136-4.902-40.802-4.772c-16.84.173-31.505 3.44-41.975 9.973L305.914 435.09c11.447-18.345 12.853-49.592 5.2-82.776c-4.05-17.564-10.502-35.45-18.5-51.84l115.427-115.43l20.274 20.274l60.817-20.273L326.954 22.87zM159.207 313.84L22.87 489.13l175.29-136.337l-38.953-38.953z"/>`),
		g.Group(children),
	)
}