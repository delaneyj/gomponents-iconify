package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 3a5 5 0 1 0 0 10a5 5 0 0 0 0-10ZM9.009 7.641a7 7 0 1 1 7.35 7.35l-9.224 7.282a1.173 1.173 0 0 1-.115.087a1.96 1.96 0 0 1-.295.167a2.322 2.322 0 0 1-1.001.207c-.851-.01-1.836-.433-2.93-1.527c-1.095-1.093-1.52-2.079-1.528-2.93a2.324 2.324 0 0 1 .207-1.002a1.974 1.974 0 0 1 .254-.41l7.282-9.224Zm.4 2.722l-6.122 7.754a.355.355 0 0 0-.021.139c.001.16.077.674.94 1.537c.865.863 1.38.94 1.54.941a.354.354 0 0 0 .137-.02l7.754-6.123a7.02 7.02 0 0 1-4.228-4.228ZM10.414 15L8 17.414L6.586 16L9 13.586L10.414 15Z"/>`),
		g.Group(children),
	)
}