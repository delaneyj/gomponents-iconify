package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageWithCurl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.5 4h18v21h-11a1 1 0 0 0-1 1h-.03c-.095.751-.487 1.438-1.05 2H8.5a3 3 0 0 1-3-3V4Zm20 21V3a1 1 0 0 0-1-1h-20a1 1 0 0 0-1 1v22a5.002 5.002 0 0 0 3.305 4.705c.361.18.925.295 1.195.295h17.5v-.03A4 4 0 0 0 29 26a1 1 0 0 0-1-1h-2.5Zm-17-15a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1h-12Zm0 3a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1h-12ZM8 16.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 0 1h-12a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1h-7Z"/>`),
		g.Group(children),
	)
}