package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emojiangry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm-73-662l-195-98q-9-9-21.5-9t-21.5 9t-9 22t9 21l173 77l-173 76q-9 9-9 21.5t9 21.5t21.5 9t21.5-9l195-98q9-9 9-21.5t-9-21.5zM256 805q0 11 7.5 19t18.5 8h460q11 0 18.5-8t7.5-19q-31-93-91-129t-165-36t-165 36t-91 129zm567-498q9-9 9-21.5t-9-21.5t-21.5-9t-21.5 9l-195 98q-9 9-9 21.5t9 21.5l195 98q9 9 21.5 9t21.5-9t9-21.5t-9-21.5l-173-76z"/>`),
		g.Group(children),
	)
}