package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditCropArtboardCropDesignImagePicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 4.5h1.5a1 1 0 0 1 1 1V7"/><path d="M4.5.5v8a1 1 0 0 0 1 1h8m-9-5h-4m9 5v4"/></g>`),
		g.Group(children),
	)
}