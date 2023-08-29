package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoginOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.25a.75.75 0 0 0 0 1.5h6A1.75 1.75 0 0 0 19.75 18V6A1.75 1.75 0 0 0 18 4.25h-6a.75.75 0 0 0 0 1.5h6a.25.25 0 0 1 .25.25v12a.25.25 0 0 1-.25.25h-6Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.25 13.115c0 .69.56 1.25 1.25 1.25h4.613l.02.22l.054.556a1.227 1.227 0 0 0 1.752.988c1.634-.783 3.212-1.958 4.466-3.267a1.253 1.253 0 0 0 0-1.734l-.1-.104a15.06 15.06 0 0 0-4.366-3.163a1.227 1.227 0 0 0-1.752.988l-.054.555l-.02.22H4.5c-.69 0-1.25.56-1.25 1.25v2.241Zm7.303.416a.75.75 0 0 0-.745-.666H4.75v-1.74h5.058a.75.75 0 0 0 .748-.704c.019-.29.042-.581.07-.871l.016-.162a13.556 13.556 0 0 1 3.516 2.607a13.568 13.568 0 0 1-3.516 2.607l-.016-.162a25.457 25.457 0 0 1-.073-.91Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}