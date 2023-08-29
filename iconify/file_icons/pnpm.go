package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pnpm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M502 10v140H362V10h140m-176 0v140H186V10h140m-176 0v140H10V10h140m352 176v140H362V186h140M512 0H352v160h160V0zM336 0H176v160h160V0zM160 0H0v160h160V0zm352 176H352v160h160V176zM336 336H176V176h160v160zm0 16H176v160h160V352zm176 0H352v160h160V352zm-352 0H0v160h160V352z"/>`),
		g.Group(children),
	)
}