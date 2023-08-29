package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 80a8 8 0 0 1 8-8h96a8 8 0 0 1 0 16h-96a8 8 0 0 1-8-8Zm104 40h-96a8 8 0 0 0 0 16h96a8 8 0 0 0 0-16Zm0 48h-72a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm-96.25 22a8 8 0 0 1-5.76 9.74a7.55 7.55 0 0 1-2 .26a8 8 0 0 1-7.75-6c-6.16-23.94-30.34-42-56.25-42s-50.09 18.05-56.25 42a8 8 0 0 1-15.5-4c5.59-21.71 21.84-39.29 42.46-48a48 48 0 1 1 58.58 0c20.63 8.71 36.88 26.29 42.47 48ZM80 136a32 32 0 1 0-32-32a32 32 0 0 0 32 32Z"/>`),
		g.Group(children),
	)
}