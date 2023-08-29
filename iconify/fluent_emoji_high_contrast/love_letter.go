package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoveLetter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.55 10.51C14.93 9.92 13.693 9 12 9c-2.5 0-4 2-4 5c0 4 6.5 9 8 9s8-5 8-9c0-3-1.5-5-4-5c-1.694 0-2.93.919-3.55 1.51a.675.675 0 0 1-.9 0Z"/><path d="M1 8a3 3 0 0 1 3-3h24a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8Zm3-1a.999.999 0 0 0-.816.421l4.61 3.073a6.504 6.504 0 0 0-.663 1.962L3 9.702V22.36l5.874-3.738a17.06 17.06 0 0 0 1.31 1.537l-6.982 4.443A.999.999 0 0 0 4 25h24a.999.999 0 0 0 .798-.398l-6.983-4.443c.463-.483.908-1 1.311-1.537L29 22.36V9.702l-4.13 2.754a6.505 6.505 0 0 0-.664-1.962l4.61-3.073A.999.999 0 0 0 28 7H4Z"/></g>`),
		g.Group(children),
	)
}