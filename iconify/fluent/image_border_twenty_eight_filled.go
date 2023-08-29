package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBorderTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.586 14.586L8 19.172V8h12v11.172l-4.586-4.586a2 2 0 0 0-2.828 0ZM17 9.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-2.646 6.146L18.707 20H9.293l4.353-4.354a.5.5 0 0 1 .708 0ZM3 6.75A3.75 3.75 0 0 1 6.75 3h14.5A3.75 3.75 0 0 1 25 6.75v14.5A3.75 3.75 0 0 1 21.25 25H6.75A3.75 3.75 0 0 1 3 21.25V6.75Zm4.25-.25a.75.75 0 0 0-.75.75v13.5c0 .414.336.75.75.75h13.5a.75.75 0 0 0 .75-.75V7.25a.75.75 0 0 0-.75-.75H7.25Z"/>`),
		g.Group(children),
	)
}