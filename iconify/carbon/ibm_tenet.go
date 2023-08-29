package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmTenet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 13h-8a1 1 0 0 0-.832 1.555l4 6a1 1 0 0 0 1.664 0l4-6A1 1 0 0 0 20 13Zm-4 5.197L13.868 15h4.264L16 18.197Z"/><path fill="currentColor" d="M16 23a7 7 0 0 1-4.42-1.58l-1.42 1.42c.03.026.057.056.088.082A9.006 9.006 0 0 0 16 25c4.963 0 9-4.038 9-9h-2c0 3.86-3.14 7-7 7Z"/><circle cx="23" cy="12" r="1" fill="currentColor"/><circle cx="9" cy="20" r="1" fill="currentColor"/><path fill="currentColor" d="M16 9c1.61 0 3.177.562 4.422 1.578L21.84 9.16c-.038-.033-.071-.07-.11-.103A9.017 9.017 0 0 0 16 7c-4.963 0-9 4.038-9 9h2c0-3.86 3.14-7 7-7Z"/><path fill="currentColor" d="M16 30C8.28 30 2 23.72 2 16S8.28 2 16 2s14 6.28 14 14s-6.28 14-14 14Zm0-26C9.383 4 4 9.383 4 16s5.383 12 12 12s12-5.383 12-12S22.617 4 16 4Z"/>`),
		g.Group(children),
	)
}