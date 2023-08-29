package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m16 .776l.027-.803l-.803.027l-1.309.046A10.75 10.75 0 0 0 5.753 4.25H3.667A2.75 2.75 0 0 0 .962 6.504l-.8 4.36L0 11.75h2.69l1.56 1.56V16l.885-.162l4.36-.8a2.75 2.75 0 0 0 2.255-2.705v-2.086a10.75 10.75 0 0 0 4.204-8.162L16 .775ZM9.348 9.988l-4.2 2.1l-1.235-1.236l2.1-4.2a9.25 9.25 0 0 1 7.954-5.107l.506-.018l-.018.506a9.25 9.25 0 0 1-5.107 7.955ZM5.75 14.2v-.736l4.268-2.135l.232-.116v1.12a1.25 1.25 0 0 1-1.025 1.23L5.75 14.2Zm-3.214-3.95l2.135-4.268l.115-.232h-1.12a1.25 1.25 0 0 0-1.229 1.025L1.8 10.25h.736ZM10.5 6a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM12 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}