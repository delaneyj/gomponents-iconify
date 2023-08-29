package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func International(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" d="M6 30.986c2.632 1.055 4.527 1.055 5.684 0c1.736-1.584.238-6.388 2.67-7.713c2.43-1.325 6.135 4.548 9.597 2.616c3.462-1.933-.326-7.087 2.076-9.176c2.403-2.09 5.527.267 6.073-3.227c.546-3.493-2.548-1.978-3.142-5.28c-.395-2.2-.395-3.357 0-3.47m.062 38.614c-1.873-1.918-2.548-3.7-2.023-5.348c.787-2.472 2.086-2.326 2.652-3.854c.566-1.528-1.033-3.705 2.515-5.565c2.366-1.24 5.619.196 9.759 4.311"/></g>`),
		g.Group(children),
	)
}