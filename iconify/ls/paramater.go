package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paramater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M414 74v37h303v60H414v37c0 11-10 20-21 20h-69c-11 0-21-9-21-20v-37H0v-60h303V74c0-11 10-20 21-20h69c11 0 21 9 21 20zM125 260h71c11 0 19 9 19 20v37h502v59H215v37c0 11-8 21-19 21h-71c-11 0-21-10-21-21v-37H0v-59h104v-37c0-11 10-20 21-20zm389 206h71c11 0 20 9 20 20v37h112v59H605v37c0 11-9 21-20 21h-71c-11 0-20-10-20-21v-37H0v-59h494v-37c0-11 9-20 20-20z"/>`),
		g.Group(children),
	)
}