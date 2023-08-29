package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LifeBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M159.4 26.23c-51.4.6-79.6 56.3-79.3 86.97c1.5 47.3 34.2 79.4 74.8 114.8c35.4 30.8 76.1 63.2 100.9 110c.1-.1.1-.2.2-.3c.1.1.1.2.2.3c24.8-46.8 65.5-79.2 100.9-110c40.6-35.4 73.3-67.5 74.8-114.8c.3-30.67-27.9-86.37-79.3-86.97c-38-.5-82.6 25.7-96.6 67.7c-14-42-58.6-68.2-96.6-67.7zM23 375v114h466V375H23zm18 18h430v78H334v-60H41v-18z"/>`),
		g.Group(children),
	)
}