package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noteslistalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024H128q-53 0-90.5-37.5T0 896V256q0-53 37.5-90.5T128 128h64V64q0-27 19-45.5T256.5 0t45 18.5T320 64v64h128V64q0-27 19-45.5T512.5 0t45 18.5T576 64v64h128V64q0-27 19-45.5T768.5 0t45 18.5T832 64v64h64q53 0 90.5 37.5T1024 256v640q0 53-37.5 90.5T896 1024zm-96-640H224q-13 0-22.5 9.5T192 416t9.5 22.5T224 448h576q13 0 22.5-9.5T832 416t-9.5-22.5T800 384zm0 192H224q-13 0-22.5 9.5T192 608t9.5 22.5T224 640h576q13 0 22.5-9.5T832 608t-9.5-22.5T800 576zm0 192H224q-13 0-22.5 9.5T192 800t9.5 22.5T224 832h576q13 0 22.5-9.5T832 800t-9.5-22.5T800 768z"/>`),
		g.Group(children),
	)
}