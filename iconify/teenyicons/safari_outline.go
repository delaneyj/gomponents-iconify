package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafariOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5 11.5l-.429-.257a.5.5 0 0 0 .686.686L3.5 11.5Zm3-5l-.257-.429l-.107.065l-.065.107l.429.257Zm5-3l.429.257a.5.5 0 0 0-.686-.686l.257.429Zm-3 5l.257.429l.107-.065l.065-.107L8.5 8.5Zm5.5-1A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM3.929 11.757l3-5l-.858-.514l-3 5l.858.514ZM6.757 6.93l5-3l-.514-.858l-5 3l.514.858Zm4.314-3.686l-3 5l.858.514l3-5l-.858-.514ZM8.243 8.07l-5 3l.514.858l5-3l-.514-.858Z"/>`),
		g.Group(children),
	)
}