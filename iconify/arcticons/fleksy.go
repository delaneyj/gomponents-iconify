package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fleksy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="21.5"/><circle cx="24" cy="24" r="21.5"/><circle cx="24" cy="24" r="21.5"/><circle cx="24" cy="24" r="21.5"/><circle cx="24" cy="13.704" r="5.704"/><circle cx="34.296" cy="24" r="5.704"/><circle cx="24" cy="34.296" r="5.704"/><circle cx="13.704" cy="24" r="5.704"/><path d="M29.825 20.477A5.703 5.703 0 0 0 28.593 24a5.704 5.704 0 0 0 1.232 3.522l2.503-2.503a1.438 1.438 0 0 0 0-2.038l-2.503-2.503ZM24 28.593a5.703 5.703 0 0 0-3.522 1.232l2.503 2.503a1.438 1.438 0 0 0 2.038 0l2.51-2.51A5.704 5.704 0 0 0 24 28.594Zm-5.819-8.122l-2.51 2.51a1.438 1.438 0 0 0 0 2.038l2.504 2.503A5.704 5.704 0 0 0 19.408 24a5.704 5.704 0 0 0-1.227-3.529ZM24 15.248c-.369 0-.737.141-1.02.424l-2.503 2.503A5.703 5.703 0 0 0 24 19.407a5.703 5.703 0 0 0 3.522-1.232l-2.503-2.503A1.437 1.437 0 0 0 24 15.248Z"/></g>`),
		g.Group(children),
	)
}