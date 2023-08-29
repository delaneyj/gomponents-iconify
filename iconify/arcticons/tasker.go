package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tasker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.89 3.5a1 1 0 0 0-1 .86l-.77 5.43a15.36 15.36 0 0 0-3.46 2L9.54 9.75a1 1 0 0 0-1.25.44l-4.1 7.11a1 1 0 0 0 .25 1.31L8.76 22a16.66 16.66 0 0 0-.14 2a16.52 16.52 0 0 0 .14 2l-4.32 3.39a1 1 0 0 0-.25 1.31l4.1 7.11a1 1 0 0 0 1.25.44l4.82-2l6.07-10.07L12 20.12l19-9.78c-.38-.2-.76-.38-1.16-.55l-.77-5.43a1.05 1.05 0 0 0-1-.86Zm18.92 6.18a1 1 0 0 0-.35.07l-4.82 1.95l-6.07 10.07L36 27.88l-19 9.78c.38.2.76.38 1.16.55l.77 5.43a1 1 0 0 0 1 .86h8.2a1 1 0 0 0 1-.86l.77-5.43a15.36 15.36 0 0 0 3.46-2l5.11 2.06a1 1 0 0 0 1.25-.44l4.1-7.11a1 1 0 0 0-.25-1.31L39.23 26a16.52 16.52 0 0 0 .14-2a16.52 16.52 0 0 0-.14-2l4.33-3.39a1 1 0 0 0 .25-1.31l-4.1-7.11a1 1 0 0 0-.9-.51Z"/>`),
		g.Group(children),
	)
}