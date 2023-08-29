package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneMilitaryTech(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 11.07l2-1.2V4h-2zM9 4v5.87l2 1.2V4z" opacity=".3"/><path fill="currentColor" d="M17 10.43V2H7v8.43c0 .35.18.68.49.86l4.18 2.51l-.99 2.34l-3.41.29l2.59 2.24L9.07 22L12 20.23L14.93 22l-.78-3.33l2.59-2.24l-3.41-.29l-.99-2.34l4.18-2.51c.3-.18.48-.5.48-.86zm-6 .64l-2-1.2V4h2v7.07zm4-1.2l-2 1.2V4h2v5.87z"/>`),
		g.Group(children),
	)
}