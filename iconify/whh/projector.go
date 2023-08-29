package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 576h-64q0 26-18.5 45t-45 19t-45.5-19t-19-45H320q0 26-18.5 45t-45 19t-45.5-19t-19-45h-64q-53 0-90.5-37.5T0 448V256q0-53 37.5-90.5T128 128h22q27-58 81.5-93T352 0t120.5 35t81.5 93h342q53 0 90.5 37.5T1024 256v192q0 53-37.5 90.5T896 576zM352 64q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zm384 192q-13 0-22.5 9.5T704 288t9.5 22.5T736 320t22.5-9.5T768 288t-9.5-22.5T736 256zm128 0q-13 0-22.5 9.5T832 288t9.5 22.5T864 320t22.5-9.5T896 288t-9.5-22.5T864 256z"/>`),
		g.Group(children),
	)
}