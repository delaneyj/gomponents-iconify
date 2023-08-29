package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m2.5 12.399l.37-.336l-.006-.007l-.007-.007l-.357.35Zm1 1.101v.5H4v-.5h-.5Zm3.5.982l.018-.5l-.053 1l.035-.5ZM7.5 7.5H7a.5.5 0 0 0 .146.354L7.5 7.5Zm6.5 0A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM2.857 12.049A6.477 6.477 0 0 1 1 7.5H0c0 2.043.818 3.897 2.143 5.249l.714-.7Zm-.727.686l1 1.101l.74-.672l-1-1.101l-.74.672ZM7.5 14a6.62 6.62 0 0 1-.465-.016l-.07.997c.177.013.355.019.535.019v-1Zm.018 0l-.5-.017l-.036 1l.5.017l.036-1ZM7 3v4.5h1V3H7Zm.146 4.854l3 3l.708-.708l-3-3l-.708.708ZM0 14h3.5v-1H0v1Zm4-.5V10H3v3.5h1Z"/>`),
		g.Group(children),
	)
}