package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundRecordingCopyright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><circle cx="36" cy="36" r="26.68" stroke-width="4.74"/><path stroke-miterlimit="10" stroke-width="8.158" d="M29.2 50.283V21.73h10.887a7.111 7.111 0 1 1 0 14.223H29.2" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}