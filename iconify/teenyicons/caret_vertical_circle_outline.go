package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretVerticalCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m10.312 9.39l.39-.312l-.624-.78l-.39.312l.624.78ZM7.5 11l-.312.39l.312.25l.312-.25L7.5 11ZM5.312 8.61l-.39-.313l-.625.781l.39.312l.625-.78Zm-.624-3l-.39.312l.624.78l.39-.312l-.624-.78ZM7.5 4l.312-.39l-.312-.25l-.312.25L7.5 4Zm2.188 2.39l.39.313l.625-.781l-.39-.312l-.625.78Zm0 2.22l-2.5 2l.624.78l2.5-2l-.624-.78Zm-1.876 2l-2.5-2l-.624.78l2.5 2l.624-.78Zm-2.5-4.22l2.5-2l-.624-.78l-2.5 2l.624.78Zm1.876-2l2.5 2l.624-.78l-2.5-2l-.624.78ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`),
		g.Group(children),
	)
}