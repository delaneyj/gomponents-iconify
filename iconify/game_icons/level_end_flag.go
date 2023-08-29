package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelEndFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 25c-8.4 0-15 6.61-15 15s6.6 15 15 15s15-6.61 15-15s-6.6-15-15-15zm7 50.53L373.1 96l81.9 20.5V75.53zM439 131v292h18V135.5l-18-4.5zm-158 54v46h46v-46h-46zm-64 64v46h46v-46h-46zm64 0v46h46v-46h-46zm-128 64v46h46v-46h-46zm64 0v46h46v-46h-46zm64 0v46h46v-46h-46zM89 377v46h46v-46H89zm64 0v46h46v-46h-46zm64 0v46h46v-46h-46zm64 0v46h46v-46h-46zM25 441v46h46v-46H25zm64 0v46h46v-46H89zm64 0v46h46v-46h-46zm64 0v46h46v-46h-46zm64 0v46h46v-46h-46zm144 0v46h46v-46h-46z"/>`),
		g.Group(children),
	)
}