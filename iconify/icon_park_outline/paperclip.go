package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26.122 37.435l14.142-14.142c2.828-2.828 4.243-9.9-.707-14.85c-4.95-4.949-12.02-3.535-14.85-.706L5.617 26.829c-1.414 1.414-3.536 6.364.707 10.606c4.242 4.243 9.192 2.121 10.607.707l18.384-18.385c1.414-1.414 2.122-4.95 0-7.07c-2.121-2.122-5.657-1.415-7.07 0L14.807 26.121"/>`),
		g.Group(children),
	)
}