package mingcute

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarWindowLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path d="M24 0v24H0V0h24ZM12.594 23.258l-.012.002l-.071.035l-.02.004l-.014-.004l-.071-.036c-.01-.003-.019 0-.024.006l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.016-.018Zm.264-.113l-.014.002l-.184.093l-.01.01l-.003.011l.018.43l.005.012l.008.008l.201.092c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022Zm-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.003-.011l.018-.43l-.003-.012l-.01-.01l-.184-.092Z"/><path fill="currentColor" fill-rule="nonzero" d="M15 14a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2h-1Z"/><path fill="currentColor" d="M7.4 5.8A7 7 0 0 1 13 3h5a3 3 0 0 1 3 3v13a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5.667a5 5 0 0 1 1-3L7.4 5.8ZM13 5a5 5 0 0 0-4 2l-3 4h13V6a1 1 0 0 0-1-1h-5Zm-8 8.333A3 3 0 0 1 5.019 13H19v6H5v-5.667Z"/></g>`),
		g.Group(children),
	)
}