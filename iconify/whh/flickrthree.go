package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickrthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800 504q-93 0-158.5-65.5T576 280t65.5-158.5T800 56t158.5 65.5T1024 280t-65.5 158.5T800 504zm0-384q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zM224 504q-93 0-158.5-65.5T0 280t65.5-158.5T224 56t158.5 65.5T448 280t-65.5 158.5T224 504z"/>`),
		g.Group(children),
	)
}