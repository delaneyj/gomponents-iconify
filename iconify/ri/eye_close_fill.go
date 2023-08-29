package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCloseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.13 15.842l-.787 2.94l-1.932-.517l.787-2.94a10.988 10.988 0 0 1-3.237-1.871l-2.153 2.153l-1.414-1.414l2.153-2.154a10.957 10.957 0 0 1-2.371-5.07l.9-.164A16.923 16.923 0 0 0 12 10c3.704 0 7.132-1.184 9.924-3.195l.9.163a10.958 10.958 0 0 1-2.37 5.071l2.153 2.154l-1.414 1.414l-2.154-2.153a10.989 10.989 0 0 1-3.237 1.872l.788 2.939l-1.932.517l-.788-2.94a11.078 11.078 0 0 1-3.74 0Z"/>`),
		g.Group(children),
	)
}