package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterGallon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M225 25v30h62V25h-62zm8 48v28.6l-5 2.5c-17 8.5-40.6 16.3-59.4 27.6c-9.6 5.8-17.6 12-23.2 19.3h221.2c-5.6-7.3-13.6-13.5-23.2-19.3c-18.8-11.3-42.4-19.1-59.4-27.6l-5-2.5V73h-46zm-112 96v16h270v-16H121zm16 34v28h238v-28H137zm-16 46v30h270v-30H121zm16 48v94h238v-94H137zm0 112v39c0 1 1.1 4.9 4 9.3c2.9 4.3 7.4 9.3 12.8 13.8c10.8 9 25.2 15.9 38.2 15.9h128c13 0 27.4-6.9 38.2-15.9c5.4-4.5 9.9-9.5 12.8-13.8c2.9-4.4 4-8.3 4-9.3v-39H137z"/>`),
		g.Group(children),
	)
}