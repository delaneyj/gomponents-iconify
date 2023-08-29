package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 7.238A2.992 2.992 0 0 1 13 5c0-1.66 1.34-3 3-3s3 1.34 3 3c0 .89-.385 1.69-1 2.238V10h3.5c.83 0 1.5.67 1.5 1.5s-.67 1.5-1.5 1.5H18v13.776A9.02 9.02 0 0 0 24.06 22h-.04c-.87 0-1.34-1.02-.77-1.68l3.39-3.96c.57-.67 1.66-.35 1.78.52l.57 3.96c.09.61-.39 1.16-1.01 1.16h-.67c-1.64 4.66-6.08 8-11.31 8s-9.67-3.34-11.31-8h-.67c-.62 0-1.1-.55-1.01-1.16l.57-3.96c.12-.87 1.21-1.19 1.78-.52l3.39 3.96c.57.66.1 1.68-.77 1.68h-.04A9.022 9.022 0 0 0 14 26.776V13h-3.5c-.83 0-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5H14V7.238ZM16 4c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1Z"/>`),
		g.Group(children),
	)
}