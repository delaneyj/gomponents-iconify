package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freecodecamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.721 36.122c-3.565-2.282-4.065-4.492-4.065-7.274c0-4.207 7.25-11.695 7.25-15.213a7.11 7.11 0 0 0-.309-2.14c2.267 1.652 5.349 4.565 5.349 8.844a8.634 8.634 0 0 1-.785 4.165a3.165 3.165 0 0 0 2.758-2.264c1.961 2.02 2.425 3.803 2.425 6.608a9.76 9.76 0 0 1-4.422 7.274a5.375 5.375 0 0 0 1.664-4.017c0-1.64-1.331-1.379-2.02-1.379a2.887 2.887 0 0 1-2.757-3.209c0-2.852.808-3.993-.357-6.228c-1.236 3.542-5.562 6.323-5.562 11.077a7.806 7.806 0 0 0 .831 3.756Zm17.802 1.919a19.483 19.483 0 0 0 0-28.082m-27.046 0a19.483 19.483 0 0 0 0 28.082"/>`),
		g.Group(children),
	)
}