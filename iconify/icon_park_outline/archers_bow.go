package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchersBow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m40.85 43.92l-1.574-.314c-2.35-.47-3.42-3.607-2.673-5.884c1.954-5.96.522-13.61-.27-16.913c-.246-1.024-1.058-1.794-2.08-2.05l-2.263-.565a3 3 0 0 1-2.183-2.183l-.565-2.262c-.256-1.023-1.025-1.834-2.05-2.08c-3.302-.792-10.954-2.225-16.913-.27c-2.276.747-5.413-.324-5.883-2.674L4.08 7.151M6 11l32 32m-26-7L43 5"/>`),
		g.Group(children),
	)
}