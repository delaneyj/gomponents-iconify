package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapanesePassingGradeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M14.52 6.141a1 1 0 0 1 .933-.641h1.094a1 1 0 0 1 .892.548c.737 1.456 3.099 4.765 6.878 6.171a1 1 0 1 1-.697 1.875c-.59-.22-1.149-.477-1.676-.763A1 1 0 0 1 21 14H11a1 1 0 0 1-.944-.669c-.52.275-1.075.53-1.663.758a1 1 0 0 1-.723-1.865c4.233-1.64 6.37-4.836 6.85-6.083Zm1.51 1.484C15.375 8.85 14.108 10.53 12.155 12h7.765a16.632 16.632 0 0 1-3.89-4.375ZM7 17a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0H9a1 1 0 1 1-2 0v-8Zm2 1v5h14v-5H9Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}