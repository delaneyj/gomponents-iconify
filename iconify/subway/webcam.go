package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webcam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248 426.7c-58.9 0-112.6-21.9-153.8-57.7L56 512h384l-38.2-143.1c-41.1 35.9-94.9 57.8-153.8 57.8zm0-149.4c47.1 0 85.3-38.2 85.3-85.3s-38.2-85.3-85.3-85.3s-85.3 38.2-85.3 85.3s38.2 85.3 85.3 85.3zm-21.3-128c11.8 0 21.3 9.5 21.3 21.3s-9.6 21.3-21.3 21.3c-11.8 0-21.3-9.5-21.3-21.3c0-11.7 9.5-21.3 21.3-21.3zM248 384c106.1 0 192-86 192-192S354.1 0 248 0S56 86 56 192s86 192 192 192zm0-320c70.7 0 128 57.3 128 128s-57.3 128-128 128s-128-57.3-128-128S177.3 64 248 64z"/>`),
		g.Group(children),
	)
}