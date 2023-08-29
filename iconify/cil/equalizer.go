package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 146.025V16H64v130.025a64.009 64.009 0 0 0 0 123.95V496h32V269.975a64.009 64.009 0 0 0 0-123.95ZM80 240a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32Zm192 50.025V16h-32v274.025a64.009 64.009 0 0 0 0 123.95V496h32v-82.025a64.009 64.009 0 0 0 0-123.95ZM256 384a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32ZM448 82.025V16h-32v66.025a64.009 64.009 0 0 0 0 123.95V496h32V205.975a64.009 64.009 0 0 0 0-123.95ZM432 176a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}