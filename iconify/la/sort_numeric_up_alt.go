package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumericUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.5 5C6.58 5 5 6.58 5 8.5V9h2v-.5C7 7.62 7.62 7 8.5 7h1c.88 0 1.5.62 1.5 1.5c0 .46-.35.98-.94 1.34c-1.23.76-2.31 1.25-3.22 1.75c-.45.26-.86.5-1.22.88c-.35.38-.62.95-.62 1.53v1h8v-2H8.44c.73-.38 1.58-.76 2.68-1.44C12.14 10.93 13 9.84 13 8.5C13 6.58 11.42 5 9.5 5h-1zm14.5.5l-.72.69L18 10.5l1.41 1.41L22 9.31V28h2V9.31l2.59 2.6L28 10.5l-4.28-4.31L23 5.5zM8.59 17l-.15.78s-.17.58-.56 1.16C7.48 19.52 6.98 20 6 20v2c1.38 0 2.32-.68 3-1.41V27h2V17H8.59z"/>`),
		g.Group(children),
	)
}