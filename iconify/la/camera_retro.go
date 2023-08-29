package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.5 6l-.313.406L10 8H9V7H5v1H3v18h26V8h-7l-1.188-1.594L20.5 6zm1 2h7l1.188 1.594L21 10h6v4h-5.813c-1.042-1.785-2.98-3-5.187-3c-2.207 0-4.145 1.215-5.188 3H5v-4h6l.313-.406zM23 11v2h2v-2zm-7 2c2.223 0 4 1.777 4 4s-1.777 4-4 4s-4-1.777-4-4s1.777-4 4-4zM5 16h5.094c-.055.32-.094.664-.094 1c0 3.3 2.7 6 6 6s6-2.7 6-6c0-.336-.04-.68-.094-1H27v8H5z"/>`),
		g.Group(children),
	)
}