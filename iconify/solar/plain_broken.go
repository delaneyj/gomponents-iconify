package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlainBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M16.212 8.848a.75.75 0 0 0-1.055-1.066l1.055 1.066Zm-5.55 5.488l5.55-5.488l-1.055-1.066l-5.55 5.488l1.056 1.066Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M18.636 15.67c-1.21 3.63-1.816 5.446-2.703 5.962c-.844.49-1.887.49-2.73 0c-.888-.516-1.493-2.331-2.704-5.962c-.194-.583-.291-.874-.454-1.118a2.172 2.172 0 0 0-.597-.597c-.244-.163-.535-.26-1.118-.454c-3.63-1.21-5.446-1.816-5.962-2.703a2.717 2.717 0 0 1 0-2.731C2.884 7.18 4.7 6.575 8.33 5.364M20.026 11.5l.326-.98c1.5-4.498 2.25-6.747 1.062-7.934c-1.187-1.187-3.436-.438-7.935 1.062L12.423 4"/></g>`),
		g.Group(children),
	)
}