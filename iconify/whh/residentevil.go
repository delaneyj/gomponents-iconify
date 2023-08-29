package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Residentevil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024H320L0 704V320L320 0h384l320 320v384zM448 512L64 335v354zM335 960h354L512 576zm177-512L689 64H335zm64 64l384 177V335z"/>`),
		g.Group(children),
	)
}