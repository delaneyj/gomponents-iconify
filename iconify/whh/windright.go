package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.888 1024q-27 0-45.5-19t-18.5-45V345l-256 153v14l-19-3q-15 6-30-4l-548-79q-18-2-30.5-15.5T.888 379V133q0-18 12.5-32t30.5-15l548-79q15-10 30-4l19-3v14l256 153V64q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v896q0 26-18.5 45t-45.5 19zm-832-868l-31 4q-14 2-23.5 12t-9.5 24v120q0 14 9.5 24t23.5 12l31 4V156zm256-34l-128 17v234l128 17V122zm192-26l-64 9v302l64 9V96zm64-6v332l256-152v-28z"/>`),
		g.Group(children),
	)
}