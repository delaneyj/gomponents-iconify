package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M41.072 5.994L3.31 16.52l9.075 9.294l8.414.146l9.683-9.44c-.256-.525-.384-.964-.384-1.318c0-.794.311-1.422.796-1.868c.83-.763 1.827-.877 2.994-.342l7.183-6.997Zm1.03.734L31.578 44.49l-9.294-9.075L22.137 27l9.375-9.518a2.542 2.542 0 0 0 1.664.495c.902-.05 1.485-.596 1.759-.917a2.35 2.35 0 0 0 .567-1.649a2.565 2.565 0 0 0-.52-1.464l7.12-7.219Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}