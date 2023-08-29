package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseSymbolForBeginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.249 1.204C5.262.636 4 1.34 4 2.501l.01 20.8c0 .716.38 1.375.998 1.735l.005.002l9.88 5.66l.006.004c.69.39 1.537.396 2.23-.005h.002l9.866-5.659l.005-.003A2.002 2.002 0 0 0 28 23.301V2.5c0-1.16-1.262-1.865-2.249-1.297L16 6.836L6.25 1.205ZM6.01 23.3L6 3.37l10 5.776l.01 19.85a.269.269 0 0 1-.127-.035l-.002-.001l-9.87-5.654v-.001L6.01 23.3Z"/>`),
		g.Group(children),
	)
}