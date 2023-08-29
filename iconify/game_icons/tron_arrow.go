package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TronArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21.844 17.813v20.843l252.062 103.28l-18.03 12.814L21.843 58.844v47.5l192 78.344l-102.813 73.218l-3.968 2.844l.032 4.844l.156 19.75l223.313 88.562l-.094 20.094l-223.033-88.47l.25 30.47l.126 13.188l.593-.22v.594l221.875 90.75l-.217 45.282L491.5 454.188l-160.625-146.25l-.188 44.343l-132.312-54.124l91.03-65.72l3.94-2.842l-.064-4.875l-.092-7.657l.218.093l-1.156-88.53L21.844 17.81zM273.78 164.97l.75 55.155l-96.53 69.72l-50.53-20.69L273.78 164.97z"/>`),
		g.Group(children),
	)
}