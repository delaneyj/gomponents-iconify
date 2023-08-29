package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moviebase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.628c3.145 4.82 10.797 19.395 16.206 19.395c7.652 0 .449-30.046-6.166-30.046c-3.112 0-6.137 4.32-8.128 7.568m-3.777 6.107c-3.712 6.155-9.812 16.37-14.341 16.37c-7.652 0-.449-30.045 6.166-30.045c4.484 0 8.464 7.904 10.04 10.65"/>`),
		g.Group(children),
	)
}