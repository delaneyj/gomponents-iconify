package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.75 12a.75.75 0 0 1-.386.656l-2.282 1.268a75.748 75.748 0 0 1-14.193 6.084l-.665.208a.75.75 0 0 1-.974-.716v-5.75c0-.391.3-.716.69-.748l.228-.018A44.247 44.247 0 0 0 10.555 12a44.24 44.24 0 0 0-6.478-.992l-.135-.01a.75.75 0 0 1-.692-.748V4.5a.75.75 0 0 1 .974-.716l.665.208a75.75 75.75 0 0 1 14.193 6.084l2.282 1.268a.75.75 0 0 1 .386.656Zm-2.294 0l-1.103-.612A74.247 74.247 0 0 0 4.75 5.52v4.04c2.93.264 5.828.81 8.654 1.631l.305.089a.75.75 0 0 1-.001 1.44l-.39.113A45.749 45.749 0 0 1 4.75 14.44v4.04a74.25 74.25 0 0 0 13.603-5.868L19.456 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}