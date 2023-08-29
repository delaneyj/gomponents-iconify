package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropboxLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m668 249l356 297l-513 316l-351-281l508-332zm-508 895l351-282l513 317l-356 296l-508-331zm864 35l513-317l351 282l-508 331l-356-296zm864-598l-351 281l-513-316l356-297l508 332zm-863 662l357 295l152-99v112l-509 305l-509-305v-112l152 99l357-295z"/>`),
		g.Group(children),
	)
}