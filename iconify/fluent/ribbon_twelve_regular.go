package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 4.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM6 1a3.5 3.5 0 0 0-2 6.373V10.5a.5.5 0 0 0 .777.416L6 10.101l1.223.815A.5.5 0 0 0 8 10.5V7.373A3.5 3.5 0 0 0 6 1Zm1 6.855v1.71l-.723-.481a.5.5 0 0 0-.554 0L5 9.566v-1.71a3.5 3.5 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}