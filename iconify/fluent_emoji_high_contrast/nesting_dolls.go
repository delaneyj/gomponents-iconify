package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestingDolls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M14 8.5a.5.5 0 0 1 1 0V9a.5.5 0 0 1-1 0v-.5Zm3 0a.5.5 0 0 1 1 0V9a.5.5 0 0 1-1 0v-.5Z"/><path d="M9 14.5c-1.145 1.717-1.763 2.985-1.943 4.5H6a1 1 0 0 0-1 1v1a9 9 0 0 0 9 9h4a9 9 0 0 0 9-9v-1a1 1 0 0 0-1-1h-1.057c-.18-1.515-.798-2.783-1.943-4.5C22 13 21.5 10 21.5 9c0-2.333-.5-7-5.5-7s-5.5 4.667-5.5 7c0 1-.5 4-1.5 5.5ZM19.603 19H23v1a6 6 0 0 1-6 6h-2a6 6 0 0 1-6-6v-1h3.397c1.299-1.436 2.574-4.053 3.232-5.84c-.363.509-.957.84-1.629.84h-1.318a.182.182 0 0 1-.182-.182c0-1.004.814-1.818 1.818-1.818h3.364c1.004 0 1.818.814 1.818 1.818c0 .1-.081.182-.182.182H18a1.998 1.998 0 0 1-1.629-.84c.658 1.787 1.933 4.404 3.232 5.84ZM16 12a4 4 0 0 1-4-4a4 4 0 0 0 4-4a4 4 0 0 0 4 4a4 4 0 0 1-4 4Z"/></g>`),
		g.Group(children),
	)
}