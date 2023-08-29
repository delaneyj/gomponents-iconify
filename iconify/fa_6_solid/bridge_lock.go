package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 64c0-17.7 14.3-32 32-32h512c17.7 0 32 14.3 32 32s-14.3 32-32 32h-40v64h-8c-61.9 0-112 50.1-112 112v24.6c-9.9 5.8-18.2 14.1-23.8 24.1c-17.6-20-43.4-32.7-72.2-32.7c-53 0-96 43-96 96v64c0 17.7-14.3 32-32 32h-32c-17.7 0-32-14.3-32-32v-64c0-53-43-96-96-96V160h72V96H64c-17.7 0-32-14.3-32-32zm376 32v64h80V96h-80zm-48 64V96h-80v64h80zM152 96v64h80V96h-80zm376 144c-17.7 0-32 14.3-32 32v48h64v-48c0-17.7-14.3-32-32-32zm-80 32c0-44.2 35.8-80 80-80s80 35.8 80 80v48c17.7 0 32 14.3 32 32v128c0 17.7-14.3 32-32 32H448c-17.7 0-32-14.3-32-32V352c0-17.7 14.3-32 32-32v-48z"/>`),
		g.Group(children),
	)
}