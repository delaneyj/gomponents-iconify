package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaggle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.352 4c-.234 0-.352.117-.352.352v23.293c0 .234.117.351.352.351h2.296c.234 0 .354-.116.354-.351v-4.836l1.809-1.723l5.238 6.664c.142.165.306.25.496.25h3.172c.166 0 .26-.048.283-.143l-.066-.359l-6.91-8.588l6.626-6.412c.123-.128.08-.498-.252-.498h-3.28c-.167 0-.333.086-.499.252L13 18.975V4.352c0-.235-.117-.352-.352-.352h-2.296z"/>`),
		g.Group(children),
	)
}