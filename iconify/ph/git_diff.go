package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitDiff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 152a8 8 0 0 0-8 8v28.69L75.72 160.4A39.71 39.71 0 0 1 64 132.12V95a32 32 0 1 0-16 0v37.13a55.67 55.67 0 0 0 16.4 39.6L92.69 200H64a8 8 0 0 0 0 16h48a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8ZM40 64a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm168 97v-37.12a55.67 55.67 0 0 0-16.4-39.6L163.31 56H192a8 8 0 0 0 0-16h-48a8 8 0 0 0-8 8v48a8 8 0 0 0 16 0V67.31l28.28 28.29A39.71 39.71 0 0 1 192 123.88V161a32 32 0 1 0 16 0Zm-8 47a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/>`),
		g.Group(children),
	)
}