package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M11.9 19.52v-7.64h1.66v7.64H11.9Zm3.31 0v-7.64h1.69v7.64h-1.69Zm3.35 0v-7.64h1.62v7.64h-1.62Z"/><path d="M16 27c6.075 0 11-4.925 11-11S22.075 5 16 5S5 9.925 5 16s4.925 11 11 11Zm6.73-5.86c.11.35-.15.7-.5.71H9.78c-.35 0-.61-.34-.51-.68l.35-1.25c.06-.23.27-.39.51-.39h.11v-7.77c-.42-.28-.5-.97.02-1.28l5.4-3.26c.24-.15.54-.15.78 0l5.39 3.26c.52.31.44 1 .01 1.28v7.76h-.03c.23 0 .44.15.51.37l.41 1.25Z"/><path d="M16 1C7.716 1 1 7.716 1 16c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15ZM3 16C3 8.82 8.82 3 16 3s13 5.82 13 13s-5.82 13-13 13S3 23.18 3 16Z"/></g>`),
		g.Group(children),
	)
}