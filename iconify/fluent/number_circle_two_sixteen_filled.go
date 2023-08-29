package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleTwoSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm4.9-1.948a1.726 1.726 0 0 1 .293-.272c.202-.148.476-.28.807-.28a1 1 0 0 1 1 1c0 .202-.1.432-.349.722c-.248.29-.588.576-.971.895l-.02.016c-.361.302-.762.637-1.07.996c-.315.366-.59.823-.59 1.371a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 0-1H7.18a2.19 2.19 0 0 1 .168-.22c.249-.29.589-.575.972-.895l.017-.014c.362-.302.765-.638 1.074-.998c.314-.367.589-.825.589-1.373a2 2 0 0 0-2-2c-.607 0-1.083.243-1.395.47a2.725 2.725 0 0 0-.503.479v.001a.5.5 0 0 0 .798.602Zm0-.001v.002"/>`),
		g.Group(children),
	)
}