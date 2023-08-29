package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M22.5 18.5h-.304c-.94 0-1.877-.274-2.564-.917a10.931 10.931 0 0 1-2.061-2.624l-.82-1.459h-.5l-.822 1.459a10.922 10.922 0 0 1-1.75 2.32a4.604 4.604 0 0 1-1.919-1.062a15.292 15.292 0 0 1-2.868-3.674L7.75 10.5h-.498l-1.141 2.043a15.3 15.3 0 0 1-3.286 4.053a3.54 3.54 0 0 1-1.325.722m0 5.182v-21h21v21h-21Zm15-13a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}