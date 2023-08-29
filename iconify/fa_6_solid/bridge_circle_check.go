package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeCircleCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32c-17.7 0-32 14.3-32 32s14.3 32 32 32h40v64H32v128c53 0 96 43 96 96v64c0 17.7 14.3 32 32 32h32c17.7 0 32-14.3 32-32v-64c0-53 43-96 96-96c6.3 0 12.4.6 18.3 1.7C367.1 231.8 426.9 192 496 192c42.5 0 81.6 15.1 112 40.2V160h-72V96h40c17.7 0 32-14.3 32-32s-14.3-32-32-32H64zm424 64v64h-80V96h80zm-128 0v64h-80V96h80zm-128 0v64h-80V96h80zm408 272a144 144 0 1 0-288 0a144 144 0 1 0 288 0zm-76.7-43.3c6.2 6.2 6.2 16.4 0 22.6l-72 72c-6.2 6.2-16.4 6.2-22.6 0l-40-40c-6.2-6.2-6.2-16.4 0-22.6s16.4-6.2 22.6 0l28.7 28.7l60.7-60.7c6.2-6.2 16.4-6.2 22.6 0z"/>`),
		g.Group(children),
	)
}