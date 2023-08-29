package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asciitilde(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M508 325v-82c-42 27-70 43-84 49c-14 5-29 8-43 8c-11 0-25-3-39-7s-42-17-81-35c-57-27-101-40-134-40c-15 0-31 2-47 8c-16 5-43 19-80 39v81c37-23 65-37 83-43c17-7 34-10 48-10c10 0 22 2 34 6s41 16 85 36c56 26 100 37 132 37c14 0 30-2 46-8c15-5 43-19 80-39z"/>`),
		g.Group(children),
	)
}