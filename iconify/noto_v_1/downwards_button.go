package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownwardsButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#db4437" d="M64 88.77c-.91 0-1.75-.46-2.25-1.23l-28.4-44.2c-.53-.83-.57-1.87-.1-2.72c.47-.85 1.37-1.39 2.34-1.39H92.4c.97 0 1.87.53 2.34 1.39c.47.85.43 1.9-.1 2.72l-28.4 44.2c-.49.77-1.33 1.23-2.24 1.23z"/>`),
		g.Group(children),
	)
}