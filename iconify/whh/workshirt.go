package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workshirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V256q0-53 37.5-90.5T128 128h64l64-96Q274 0 384 0h256q110 0 128 32l64 96h64q53 0 90.5 37.5T1024 256v640q0 53-37.5 90.5T896 1024zM320 64l-64 96l96 96l96-128zm81 237q-25 19-49 19q-32 0-64-32l-64-64q-13-13-22-32h-74q-26 0-45 18.5T64 256v640q0 27 18.5 45.5T128 960h192l129-617zM704 64l-128 64l96 128l96-96zm256 192q0-27-19-45.5T896 192h-75q-8 19-21 32l-64 64q-32 32-64 32q-24 0-49-19l-49 43l130 616h192q26 0 45-18.5t19-45.5V256z"/>`),
		g.Group(children),
	)
}