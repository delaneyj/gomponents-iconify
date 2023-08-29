package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.79 4.46l.71-.71L15.26 15.5l.71.72l4.28 4.28l-.71.71l-4.35-4.36c-.92.65-2.01 1.04-3.19 1.15v3h-1v-3c-3.36-.27-6-3.08-6-6.5V11h1v.5a5.5 5.5 0 0 0 5.5 5.5c1.09 0 2.11-.32 2.97-.87l-2.24-2.24c-.23.07-.48.11-.73.11A2.5 2.5 0 0 1 9 11.5v-.84l-6.21-6.2M17 11.5V11h1v.5c0 1.5-.5 2.88-1.36 4l-.71-.74c.67-.91 1.07-2.04 1.07-3.26M11.5 3A2.5 2.5 0 0 1 14 5.5v6c0 .39-.09.76-.25 1.09l-.78-.79l.03-.3v-6A1.5 1.5 0 0 0 11.5 4A1.5 1.5 0 0 0 10 5.5v3.34l-1-1V5.5A2.5 2.5 0 0 1 11.5 3M10 11.67c.09.69.64 1.24 1.33 1.33L10 11.67Z"/>`),
		g.Group(children),
	)
}