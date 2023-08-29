package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlainThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m20.352 10.52l-1.716 5.15c-1.21 3.63-1.816 5.446-2.703 5.962c-.844.49-1.887.49-2.73 0c-.888-.516-1.493-2.331-2.704-5.962c-.194-.583-.291-.874-.454-1.118a2.172 2.172 0 0 0-.597-.597c-.244-.163-.535-.26-1.118-.454c-3.63-1.21-5.446-1.816-5.962-2.703a2.717 2.717 0 0 1 0-2.731C2.884 7.18 4.7 6.575 8.33 5.364l5.15-1.716c4.498-1.5 6.747-2.25 7.934-1.062c1.187 1.187.438 3.436-1.062 7.935Z" clip-rule="evenodd" opacity=".5"/><path d="M13.046 10.902a.75.75 0 0 1 .006-1.06l4.21-4.164a.75.75 0 0 1 1.055 1.066l-4.21 4.164a.75.75 0 0 1-1.061-.006Z"/></g>`),
		g.Group(children),
	)
}