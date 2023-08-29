package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoCodesandbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.34 6.423L12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423ZM12 3.155L9.67 4.5L12 5.845L14.33 4.5L12 3.155Zm4.33 2.5L12 8.155l-4.33-2.5L5.34 7L12 10.845L18.66 7l-2.33-1.345Zm3.33 3.077L13 12.577v7.69l2.34-1.35v-4.994l4.32-2.495V8.732Zm0 5.006l-2.32 1.34v2.684l2.32-1.34v-2.684Zm-15.32-2.31l4.32 2.495v4.994l2.34 1.35v-7.69L4.34 8.732v2.696Zm0 2.31v2.685l2.32 1.34v-2.686l-2.32-1.34Z"/>`),
		g.Group(children),
	)
}