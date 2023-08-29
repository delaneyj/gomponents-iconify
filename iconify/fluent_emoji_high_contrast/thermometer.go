package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M26.97 21.97a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path d="m8.834 3.176l12.903 12.903a7 7 0 1 1-5.658 5.658L3.176 8.834a4.004 4.004 0 0 1 0-5.658a4.004 4.004 0 0 1 5.65-.01l.008.01Zm9.48 17.968a5 5 0 1 0 2.83-2.83l-2.419-2.419l-.725.725a.996.996 0 1 1-1.41-1.41l.725-.725l-2.83-2.83l-.725.725a.996.996 0 1 1-1.41-1.41l.725-.725L10.24 7.41l-.72.72a.996.996 0 1 1-1.41-1.41L8.83 6L7.42 4.59c-.79-.78-2.05-.78-2.83 0c-.78.78-.78 2.05 0 2.83l13.724 13.724Z"/></g>`),
		g.Group(children),
	)
}