package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundwaveSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM12.75 7a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0V7Zm-5 2a.75.75 0 0 0-1.5 0v6a.75.75 0 0 0 1.5 0V9Zm10 1a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 1.5 0v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}