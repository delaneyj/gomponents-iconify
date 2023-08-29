package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFinland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M61.1 25C58 11.8 46.1 2 32 2h-1v23h30.1M17 6C10 10.1 4.7 16.9 2.8 25H17V6M2.8 39C4.7 47.1 10 53.9 17 58V39H2.8zM31 62h1c14.2 0 26-9.8 29.2-23H31v23"/><path fill="#428bc1" d="M61.1 25H31V2c-5.1.2-9.9 1.6-14 4v19H2.8c-.5 2.2-.8 4.6-.8 7s.3 4.8.8 7H17v19c4.1 2.4 8.9 3.8 14 4V39h30.2c.5-2.2.8-4.6.8-7s-.3-4.8-.9-7"/>`),
		g.Group(children),
	)
}