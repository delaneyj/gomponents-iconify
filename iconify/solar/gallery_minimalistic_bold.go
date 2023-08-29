package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GalleryMinimalisticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM16 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-9.68 3.104a1.55 1.55 0 0 1 2.184.073l2.648 2.81a2.294 2.294 0 0 0 3.042.266a1.567 1.567 0 0 1 2.02.123l2.268 2.166a.75.75 0 0 0 1.036-1.084L17.25 15.29a3.067 3.067 0 0 0-3.953-.24a.795.795 0 0 1-1.054-.093l-2.647-2.81a3.05 3.05 0 0 0-4.296-.143l-.81.752a.75.75 0 1 0 1.02 1.1l.81-.753Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}